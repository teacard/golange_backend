# 草原精靈獸 / Grass Spirit

**ID：** `creature_grassspirit`
**稀有度：** ★★☆☆（Uncommon）
**個性：** 悠閒型

---

## 圖鑑介紹

棲息於廣闊草原的溫馴精靈獸，頭頂總是插著一片嫩綠的葉子，據說那片葉子能感應天氣的變化。性格悠閒，不喜爭鬥，喜歡在草地上發呆或追著蒲公英跑。牠的腮紅不論何時都是紅的，沒有人知道是天生的還是害羞的。

遇到陌生人時會先愣住，接著用大眼睛盯著對方看很久，最後搖搖尾巴表示接受。

---

## 外觀描述

- **體型：** 中型，圓潤飽滿，四肢短小可愛
- **頭部：** 大圓頭，有一對直立耳，耳內淡粉色
- **特徵：** 頭頂插一片雙葉，葉脈清晰
- **尾巴：** 中等長度，末端微捲
- **表情：** 大眼睛有明顯高光，嘴角微微上揚

---

## 色票

| 部位 | 色碼 |
|------|------|
| 身體主色 | `#e8d8b0` |
| 身體輪廓 | `#b8a070` |
| 肚皮 | `#f5ecd8` |
| 肚皮輪廓 | `#c8b888` |
| 尾巴 / 後腳 | `#c4a882` |
| 腳掌 | `#d4c090` |
| 耳內 | `#f0c0b0` |
| 腮紅 | `rgba(240,120,110,0.35)` |
| 鼻子 | `#d4907a` |
| 頭頂葉子 | `#8aba60` / 輪廓 `#6a9a40` |

---

## 動畫參數

```yaml
idle:
  breathe_speed: 1.2
  breathe_amp: 3
  tail_speed: 0.9
  tail_amp: 0.12
  blink_interval: 0.8

walk:
  cycle: 3.0
  body_bounce: 5
  body_sway: 0.04
  leg_amp: 0.35
  tail_amp: 0.2

catch:
  shake_freq: 18
  shake_amp: 8
  bounce_amp: 10
  duration: 1.5
  blush_alpha: 0.6
```

---

## 生物屬性

```json
{
  "id": "creature_grassspirit",
  "name": "草原精靈獸",
  "name_en": "Grass Spirit",
  "rarity": "uncommon",
  "personality": "悠閒型",
  "habitat": "草原",
  "catchable_time": "全時段",
  "description": "棲息於廣闊草原的溫馴精靈獸，頭頂插著一片能感應天氣的嫩葉。",
  "appearance": {
    "body_color": "#e8d8b0",
    "body_stroke": "#b8a070",
    "special_parts": ["head_leaf", "round_ears"]
  },
  "animation": {
    "idle_breath_speed": 1.2,
    "idle_breath_amp": 3,
    "walk_cycle": 3.0,
    "walk_leg_amp": 0.35
  }
}
```

---

## SVG / Canvas 程式碼

> 完整 Canvas 程式碼見 `assets/creatures/grassspirit.html`
> 此檔為設計來源，Ebitengine 實作時依動畫參數重新繪製

**繪製重點：**
- 身體為兩個疊合橢圓（body + belly）
- 頭頂葉子用 `bezierCurveTo` 畫雙葉形
- 眼睛需含白色高光圓點（半徑 3px，偏右上）
- 尾巴用 `quadraticCurveTo` 帶弧度
