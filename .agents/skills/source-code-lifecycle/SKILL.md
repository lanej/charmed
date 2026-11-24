---
name: source-code-lifecycle
description: Use when applying installed project policy for project kind, tool choices, authority, or adoption.
---

# Installed project policy

## Project kind unspecified

Preserve authored project and directory requirements. No project-wide kind obligations are installed. Do not infer a category from files or dependencies or prescribe architecture from an unselected kind.

## Authority: owned

The project controls tools for its owned surfaces, subject to the selected adoption policy and authored project requirements. Owning a client does not authorize redesigning an external provider’s protocol.

Apply authority and adoption together. Authority is independent of lifecycle stage, project kind, and public visibility; it does not grant publication or migration authority beyond the task.

## Tool adoption: hold

Preserve the established stack, including extensions to existing surfaces. Do not use the application tool catalog to introduce a competing stack. For an unestablished capability, obtain a project choice before selecting a stack. Do not load the catalog merely to perform ordinary work on the retained stack.

Apply authority and adoption together. Adoption does not establish project agreement or waive repository requirements.

Preserve authored repository requirements. Use the installed testing, comment, and language skills for the work they cover. Apply the authority and adoption policy above when choosing application libraries, API protocols, or migration scope.

For CLI work, use [CLI construction](../cli/SKILL.md) for language tools and command discovery. Use the installed `data-pipeline` or `ui-acceptance` skill when the task calls for it. Those skills are selected explicitly during setup.
