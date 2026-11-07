// Progressive enhancement for saved terminal variant tabs.
// Without JavaScript, all snapshots remain visible and links scroll to them.
(function () {
  "use strict";

  document.querySelectorAll("[data-variant-tabs]").forEach(function (group) {
    var tablist = group.querySelector('[role="tablist"]');
    if (!tablist) return;
    var tabs = Array.from(tablist.querySelectorAll('[role="tab"]'));
    var panels = tabs.map(function (tab) {
      return document.getElementById(tab.getAttribute("aria-controls"));
    });
    if (!tabs.length || panels.some(function (panel) { return !panel; })) return;

    function activate(index, shouldFocus) {
      tabs.forEach(function (tab, i) {
        var selected = i === index;
        tab.setAttribute("aria-selected", String(selected));
        tab.tabIndex = selected ? 0 : -1;
        panels[i].hidden = !selected;
      });
      if (shouldFocus) tabs[index].focus();
    }

    var deepLinked = panels.findIndex(function (panel) {
      return "#" + panel.id === window.location.hash;
    });
    activate(deepLinked < 0 ? 0 : deepLinked, false);

    tabs.forEach(function (tab, index) {
      tab.addEventListener("click", function (event) {
        event.preventDefault();
        activate(index, false);
      });

      tab.addEventListener("keydown", function (event) {
        var next;
        switch (event.key) {
          case "ArrowRight": next = (index + 1) % tabs.length; break;
          case "ArrowLeft": next = (index - 1 + tabs.length) % tabs.length; break;
          case "Home": next = 0; break;
          case "End": next = tabs.length - 1; break;
          default: return;
        }
        event.preventDefault();
        activate(next, true);
      });
    });
  });
})();
