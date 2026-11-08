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

// The mobile menu is a native disclosure: it works even without this enhancer.
// Close it on navigation without leaving keyboard focus in hidden content.
(function () {
  "use strict";
  var menu = document.querySelector("[data-mobile-nav]");
  if (!menu) return;
  var summary = menu.querySelector("summary");

  function closeMenu(restoreFocus) {
    menu.open = false;
    if (restoreFocus) summary.focus({ preventScroll: true });
  }

  menu.addEventListener("click", function (event) {
    var link = event.target.closest("a[href]");
    if (!link || !menu.contains(link) || event.button !== 0 ||
        event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
    var href = link.getAttribute("href");
    if (!href.startsWith("#")) return; // Let page navigation work normally.
    var target;
    try { target = document.getElementById(decodeURIComponent(href.slice(1))); }
    catch (_) { return; }
    if (!target) return;
    closeMenu(false);
    if (!target.hasAttribute("tabindex")) target.setAttribute("tabindex", "-1");
    target.focus({ preventScroll: true });
    // Native fragment navigation updates the URL/history after the menu closes.
  });

  document.addEventListener("keydown", function (event) {
    if (event.key === "Escape" && menu.open) {
      event.preventDefault();
      closeMenu(true);
    }
  });
  document.addEventListener("click", function (event) {
    if (menu.open && !menu.contains(event.target)) {
      closeMenu(menu.contains(document.activeElement));
    }
  });
  window.matchMedia("(max-width: 760px)").addEventListener("change", function () {
    closeMenu(false);
  });
  window.addEventListener("pageshow", function () { closeMenu(false); });
})();
