package models

// CryptoPrice maps dynamic JSON responses like:
// {
//   "bitcoin": { "usd": 43000 }
// }
type CryptoPrice map[string]map[string]float64
