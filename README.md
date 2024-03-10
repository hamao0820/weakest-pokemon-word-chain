# 最短ポケモンしりとり

[ここ](https://hamao0820.github.io/weakest-pokemon-word-chain)で遊べます。

## 遊び方

最初のポケモンを入力してボタンを押すと、最短で終わるポケモンしりとりができます。  
同じポケモンから始めても、毎回異なるしりとりができるので、何度か試してみてください。

## しりとりのルール

しりとりのルールは[世界ultimateしりとり協会](https://w.atwiki.jp/ultimate/pages/16.html)を参考にしています。

## アルゴリズム

有向グラフを構築し、幅優先探索を用いています。  
選ばれたポケモンを始点として、「ン」で終わるポケモンが見つかった時点で探索を打ち切ります。

## データの出典

ポケモンのデータは[ポケモン王国攻略館](https://pente.koro-pokemon.com/zukan/)から取得しています。  
画像データは[PokeApi](https://pokeapi.co/)から取得しています。
