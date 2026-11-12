#!/usr/bin/env python3
"""Enforce the source and generated documentation contract, without dependencies.

Words are visible Unicode word tokens; inline code counts, link destinations do
not. A run spans headings until a real code block, image or multirow table.
"""
import argparse
from dataclasses import dataclass, field
from html.parser import HTMLParser
import json
from pathlib import Path
import re
import sys
from urllib.parse import unquote, urlsplit

ROOT = Path(__file__).resolve().parents[1]
WORD = re.compile(r"\b[\w]+(?:[’'-][\w]+)*\b", re.UNICODE)


def visible(text):
    return re.sub(r"!?\[([^\]]*)\]\([^)]*\)", r"\1", text)


def words(text):
    return len(WORD.findall(visible(text)))


@dataclass
class Finding:
    path: str
    line: int
    rule: str
    actual: object
    allowed: object
    correction: str

    def __str__(self):
        return f"{self.path}:{self.line}: {self.rule}: {self.actual}; expected {self.allowed}. {self.correction}"


def bounds(findings, path, line, rule, actual, allowed, correction):
    if not allowed[0] <= actual <= allowed[1]:
        findings.append(Finding(str(path), line, rule, actual, f"{allowed[0]}–{allowed[1]}", correction))


def markdown_blocks(source):
    """The authored GFM subset: fences, tables, images, headings, lists, prose.

    Heading and blank-line boundaries never reset an uninterrupted text run.
    """
    lines = source.splitlines()
    i = 0
    while i < len(lines):
        line, number = lines[i], i + 1
        if not line.strip():
            i += 1
            continue
        fence = re.match(r"^\s*(`{3,}|~{3,})(.*)$", line)
        if fence:
            marker = fence[1][0]
            content = []
            i += 1
            while i < len(lines) and not re.match(rf"^\s*{marker}{{3,}}\s*$", lines[i]):
                content.append(lines[i]); i += 1
            yield "code", number, "\n".join(content)
            i += 1
            continue
        if re.match(r"^#{1,6}\s", line):
            yield "heading", number, line.lstrip('# ').strip()
            i += 1
            continue
        if line.startswith('|'):
            content = []
            while i < len(lines) and lines[i].startswith('|'):
                content.append(lines[i]); i += 1
            yield "table", number, "\n".join(content)
            continue
        if re.fullmatch(r"!\[[^\]]+\]\([^)]*\)\s*", line):
            yield "image", number, line
            i += 1
            continue
        content = []
        while i < len(lines) and lines[i].strip() and not re.match(r"^(?:#{1,6}\s|```|~~~|\|)", lines[i]):
            content.append(lines[i]); i += 1
        text = " ".join(content)
        if re.match(r"^(?:[-*+] |\d+\. )", text):
            yield "list", number, text
        elif re.fullmatch(r"(?:\[[^\]]+\]\([^)]*\)[\s·]*)+", text):
            yield "links", number, text
        else:
            yield "paragraph", number, text


def check_markdown(source, path, profile, sections=(), artifact=None):
    findings, run, headings = [], [], []

    def flush():
        if not run:
            return
        count = sum(value for _, value in run)
        bounds(findings, path, run[0][0], 'run.words', count, profile['run_words'], 'Combine related explanation or replace dense prose with a useful example or comparison table.')
        bounds(findings, path, run[0][0], 'run.blocks', len(run), profile['run_blocks'], 'Combine related paragraphs; headings do not reset the budget.')
        if len(run) > 1 and count / len(run) < profile['mean_words_min']:
            findings.append(Finding(str(path), run[0][0], 'run.mean_words', round(count / len(run), 1), f">= {profile['mean_words_min']}", 'Combine fragments into coherent paragraphs instead of splitting to pass.'))
        run.clear()

    for kind, line, text in markdown_blocks(source):
        if kind == 'heading':
            headings.append(text)
        elif kind == 'links':
            run.append((line, words(text)))
        elif kind in ('paragraph', 'list'):
            if kind == 'paragraph':
                bounds(findings, path, line, 'paragraph.words', words(text), profile['paragraph_words'], 'Combine a fragment with its explanation, or rewrite the paragraph around one point.')
            run.append((line, words(text)))
        elif kind == 'code':
            bounds(findings, path, line, 'code.lines', len(text.splitlines()), profile['code_lines'], 'Keep the minimal usage here and link the complete program.')
            useful = [line for line in text.splitlines() if line.strip() and not line.lstrip().startswith(('#', '//')) and re.search(r'[\w]', line)]
            if useful:
                flush()
            else:
                findings.append(Finding(str(path), line, 'artifact.code', 'empty or comments only', 'useful content', 'Supply an actual example; decorative code blocks do not reset the budget.'))
        elif kind == 'image':
            alt = re.match(r"!\[([^\]]*)\]", text)[1]
            if artifact and artifact['image_alt_words'][0] <= words(alt) <= artifact['image_alt_words'][1]:
                flush()
            else:
                findings.append(Finding(str(path), line, 'artifact.image', words(alt), str(artifact['image_alt_words']) + ' descriptive words', 'Describe what the image demonstrates; decoration does not reset the text budget.'))
        elif kind == 'table':
            rows = text.splitlines()
            if len(rows) >= 4 and all(row.count('|') >= 3 for row in rows):
                flush()
                if artifact:
                    bounds(findings, path, line, 'table.rows', len(rows)-2, artifact['table_rows'], 'Keep one focused comparison in the table.')
                    for index, row in enumerate(rows):
                        if index == 1:
                            continue
                        # Ignore escaped pipes and pipes inside code spans.
                        cells = re.sub(r'`[^`]*`', lambda match: match[0].replace('|', '/'), row).strip('|').split('|')
                        for cell in cells:
                            bounds(findings, path, line+index, 'table.cell_words', words(cell), artifact['table_cell_words'], 'Use a concise field or contract, rather than hiding a paragraph in a cell.')
            else:
                run.append((line, words(text)))
    flush()
    for section in sections:
        if section not in headings:
            findings.append(Finding(str(path), 1, 'section.required', section, 'present', 'Add the missing task section.'))
    return findings


def check_sources(root, config):
    findings = []
    actual = {p.name for p in (root / 'docs').glob('*.md')}
    declared = set(config['pages'])
    for name in sorted(actual ^ declared):
        findings.append(Finding('docs/' + name, 1, 'profile.coverage', 'unclassified or missing page', 'one declared profile per page', 'Update docs/standards.json with the page intent.'))
    for name in sorted(actual & declared):
        spec = config['pages'][name]
        findings += check_markdown((root / 'docs' / name).read_text(), 'docs/' + name, config['profiles'][spec['profile']], spec.get('sections', []), config['artifact'])
    for path in sorted((root / 'docs/snippets').glob('*.go')):
        if path.name == 'render.go':
            continue
        source = path.read_text()
        # Count the complete consumer body plus helpers, excluding package/import scaffolding.
        source = re.sub(r'(?ms)^package [^\n]+\n|^import \(.*?^\)\n|^import "[^\n]+\n', '', source).strip()
        source = re.sub(r'(?m)^func Example[^\n]+\{\n', '', source, count=1)
        count = len(source.splitlines()) - 1
        bounds(findings, path.relative_to(root), 1, 'snippet.lines', count, config['component']['code_lines'], 'Simplify sample data and retain the complete consumer usage; do not hide dependencies.')
    guidance = json.loads((root / 'docs/component-guidance.json').read_text())
    known_ids = set(re.findall(r'case "([^"]+)":', (root / 'docs/snippets/render.go').read_text()))
    for identifier in sorted(set(guidance['use']) - known_ids):
        findings.append(Finding('docs/component-guidance.json', 1, 'guidance.unknown_id', identifier, 'registered consumer example', 'Correct the component ID rather than leaving unused guidance.'))
    for name, text in guidance['integration'].items():
        bounds(findings, 'docs/component-guidance.json', 1, f'integration.words[{name}]', words(text), config['component']['integration_words'], 'Explain updates, focus and sizing together as one useful paragraph.')
    return findings


@dataclass
class Node:
    tag: str
    attrs: dict = field(default_factory=dict)
    line: int = 1
    children: list = field(default_factory=list)

    def walk(self):
        yield self
        for child in self.children:
            if isinstance(child, Node):
                yield from child.walk()

    def text(self):
        return ''.join(c.text() if isinstance(c, Node) else c for c in self.children)

    def has(self, css_class):
        return css_class in self.attrs.get('class', '').split()


class Document(HTMLParser):
    def __init__(self, source):
        super().__init__(convert_charrefs=True)
        self.root = Node('document')
        self.stack = [self.root]
        self.feed(source)

    def handle_starttag(self, tag, attrs):
        node = Node(tag, dict(attrs), self.getpos()[0])
        self.stack[-1].children.append(node)
        if tag not in ('meta', 'link', 'img', 'br', 'hr', 'input', 'source', 'wbr'):
            self.stack.append(node)

    def handle_endtag(self, tag):
        for index in range(len(self.stack)-1, 0, -1):
            if self.stack[index].tag == tag:
                del self.stack[index:]
                break

    def handle_data(self, data):
        self.stack[-1].children.append(data)


def check_rendered(site, config):
    findings, docs, ids = [], {}, {}
    for path in sorted(site.glob('*.html')):
        docs[path.name] = Document(path.read_text()).root
        ids[path.name] = set()
        for node in docs[path.name].walk():
            identifier = node.attrs.get('id')
            if identifier:
                if identifier in ids[path.name]:
                    findings.append(Finding(path.name, node.line, 'link.duplicate_id', identifier, 'unique', 'Give the heading or artifact a stable unique ID.'))
                ids[path.name].add(identifier)
    for name, doc in docs.items():
        for node in doc.walk():
            for attribute in ('href', 'src'):
                target = node.attrs.get(attribute)
                if target is None:
                    continue
                url = urlsplit(target)
                if url.scheme or url.netloc:
                    continue
                file = unquote(url.path) or name
                if not (site / file).is_file() or (url.fragment and unquote(url.fragment) not in ids.get(file, set())):
                    findings.append(Finding(name, node.line, 'link.local', target, 'existing file and anchor', 'Link directly to the published HTML page and stable heading.'))
            if not node.has('component'):
                continue
            nodes = list(node.walk())
            for css in ('component-use', 'component-preview', 'component-source', 'component-integration', 'component-run'):
                if not any(n.has(css) for n in nodes):
                    findings.append(Finding(name, node.line, 'component.section', css, 'present', 'Use the shared component template.'))
            for n in nodes:
                if n.has('component-use'):
                    bounds(findings, name, n.line, 'component.use_words', words(n.text()), config['component']['use_words'], 'Explain the purpose in one concise paragraph.')
                if n.has('source-code'):
                    bounds(findings, name, n.line, 'component.code_lines', len(n.text().strip().splitlines()), config['component']['code_lines'], 'Keep a complete minimal usage example.')
                if n.attrs.get('id', '').endswith('-example') and n.attrs.get('role') == 'tabpanel':
                    previews = [p for p in n.walk() if p.has('terminal')]
                    for preview in previews:
                        text = preview.text().strip('\n')
                        bounds(findings, name, preview.line, 'preview.rows', len(text.splitlines()), config['component']['preview_rows'], 'Use a representative bounded allocation.')
                        visible_fill = any('background-color:' in cell.attrs.get('style', '') and cell.text() for cell in preview.walk())
                        if not text.strip() and not visible_fill:
                            findings.append(Finding(name, preview.line, 'preview.empty', 0, 'nonempty useful output', 'Initialize asynchronous models and render representative data.'))
    if not docs:
        findings.append(Finding(str(site), 1, 'site.missing', 0, 'generated pages', 'Run just docs first.'))
    return findings


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--rendered', type=Path)
    parser.add_argument('--root', type=Path, default=ROOT)
    args = parser.parse_args()
    config = json.loads((args.root / 'docs/standards.json').read_text())
    findings = check_sources(args.root, config)
    if args.rendered:
        findings += check_rendered(args.rendered, config)
    findings = apply_exceptions(findings, config['exceptions'])
    for finding in findings:
        print(finding)
    if findings:
        print(f'Documentation standard: {len(findings)} violation(s).', file=sys.stderr)
        return 1
    print('Documentation standard: source' + (' and generated HTML' if args.rendered else '') + ' passed.')
    return 0


def apply_exceptions(findings, exceptions, today=None):
    if exceptions:
        # Exceptions require exact findings, a reason, an owner and a review date;
        # there are deliberately no broad file or rule suppressions.
        from datetime import date
        today = today or date.today()
        for exception in exceptions:
            required = {'path', 'line', 'rule', 'actual', 'reason', 'owner', 'review_by'}
            if not required <= exception.keys() or not exception['owner'].strip() or not exception['reason'].strip() or date.fromisoformat(exception['review_by']) <= today:
                raise ValueError('Incomplete or expired documentation exception')
            matched = [f for f in findings if (f.path, f.line, f.rule, f.actual) == (exception['path'], exception['line'], exception['rule'], exception['actual'])]
            if len(matched) != 1:
                raise ValueError('Stale or ambiguous documentation exception')
            print('Reviewed exception: ' + str(matched[0]))
            findings.remove(matched[0])
    return findings


if __name__ == '__main__':
    sys.exit(main())
