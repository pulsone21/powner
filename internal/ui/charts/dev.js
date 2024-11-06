(() => {
  const data = JSON.parse(document.currentScript.previousSibling.node);
  const config = {
    type: "radar",
    data: data,
    options: {
      elements: {
        line: {
          borderWidth: 3,
        },
      },
    },
  };
  const canvas = document.currentScript.previousSibling.previousSibling;
  new Chart(canvas, config);
})();
