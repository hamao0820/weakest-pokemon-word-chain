(async () => {
  const res = await fetch("data/zukan.json");
  const data = await res.json();
  const $select = document.querySelector("#select");
  const noDist = {};
  data.data.forEach((p) => {
    const $option = document.createElement("option");
    $option.value = p.no;
    $option.text = p.name;
    $select.appendChild($option);
    noDist[p.no] = p.name;
  });
  const $button = document.querySelector("button");
  const $result = document.querySelector("#result");
  $button.addEventListener("click", () => {
    const path = goGetShortestChain(Number($select.value));
    if (path.length === 0) {
      $result.textContent = "このポケモンから始めると、しりとりは終了しません。";
      return;
    }
    const result = path.map((no) => noDist[no]).join(" -> ");
    $result.textContent = result;
  });
})();
