import { createHighlighterCore } from "shiki/core";
import { createJavaScriptRegexEngine } from "shiki/engine/javascript";
import go from "shiki/langs/go.mjs";
import githubDark from "shiki/themes/github-dark.mjs";

const highlighter = createHighlighterCore({
  themes: [githubDark],
  langs: [go],
  engine: createJavaScriptRegexEngine(),
});
export async function highlightGo(source: string) {
  return (await highlighter).codeToHtml(source, {
    lang: "go",
    theme: "github-dark",
  });
}
