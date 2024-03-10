const createResultItem = (name, imgSrc) => {
  const $div = document.createElement("div");
  $div.classList.add("result-item");
  const $img = document.createElement("img");
  $img.src = imgSrc;
  const $span = document.createElement("span");
  $span.textContent = name;
  $div.appendChild($img);
  $div.appendChild($span);
  return $div;
};

(async () => {
  const res = await fetch("data/zukan.json");
  const data = await res.json();
  const $select = document.querySelector("#name-select");
  const $input = document.querySelector("#no-input");
  const noDist = {};
  data.data.forEach((p) => {
    const $option = document.createElement("option");
    $option.value = p.no;
    $option.text = p.name;
    $select.appendChild($option);
    noDist[p.no] = p.name;
  });

  $input.addEventListener("change", (e) => {
    $select.value = e.target.value;
  });

  $input.addEventListener("input", (e) => {
    $select.value = e.target.value;
  });

  $input.addEventListener("blur", (e) => {
    const v = Number(e.target.value);
    if (v < 1) {
      e.target.value = 1;
    } else if (v > 1024) {
      e.target.value = 1024;
    }
    $select.value = e.target.value;
  });

  $select.addEventListener("change", (e) => {
    $input.value = e.target.value;
  });

  const $button = document.querySelector("button");
  const $result = document.querySelector("#result");
  $button.addEventListener("click", async () => {
    const path = goGetShortestChain(Number($select.value));
    if (path.length === 0) {
      $result.textContent =
        "このポケモンから始めると、しりとりは終了しません。";
      return;
    }
    const data = await Promise.all(
      path.map(async (no) => {
        const url = "https://pokeapi.co/api/v2/pokemon/" + no;
        const res = await fetch(url);
        return res.json();
      })
    );

    const sources = data.map((d) => {
      return d.sprites.front_default;
    });

    $result.innerHTML = "";
    for (let i = 0; i < path.length; i++) {
      $result.appendChild(createResultItem(noDist[path[i]], sources[i]));
    }
  });
})();
