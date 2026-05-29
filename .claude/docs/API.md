# API 端點文件

Base URL：`http://localhost:8080`

---

## 生物

### GET /api/creatures
取得所有生物列表。

**Response**
```json
{
  "data": [
    {
      "id": "creature_grassspirit",
      "rarity": "uncommon",
      "personality": "悠閒型",
      "habitat": "草原",
      "colors": {"primary": "#e8d8b0", "secondary": "#b8a070"},
      "special_parts": [{"code": "head_leaf", "color": "#8aba60"}, {"code": "round_ears"}],
      "animation": {"idle": {...}, "walk": {...}, "catch": {...}}
    }
  ]
}
```

---

### GET /api/creatures/:id
取得單一生物資料。`:id` 為 `creature_id`。

**Response**：同上，`data` 為單一物件。

**Error**
```json
{"error": "找不到該生物"}  // 404
```

---

## 玩家（待實作）

### GET /api/player/inventory
取得玩家已收集的生物清單。

### POST /api/player/catch
捕獲生物。

---

## 狀態

| Method | Path | Handler | 狀態 |
|--------|------|---------|------|
| GET | `/api/creatures` | `creatureHandler.GetAll` | ✅ |
| GET | `/api/creatures/:id` | `creatureHandler.GetByID` | ✅ |
| GET | `/api/player/inventory` | `playerHandler.GetInventory` | ⬜ |
| POST | `/api/player/catch` | `playerHandler.CatchCreature` | ⬜ |
