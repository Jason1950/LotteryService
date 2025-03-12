# LotteryService

LotteryService 是一個基於 Go 語言的簡單服務，旨在提供彩票系統的功能。此服務包含了用戶登入 (login) 及彩票投注 (bet) API，並結合了區塊鏈技術來處理彩票的投注與記錄。

## 目錄

- [簡介](#簡介)
- [安裝](#安裝)
- [API 文檔](#api-文檔)
  - [Login API](#login-api)
  - [Bet API](#bet-api)
- [使用說明](#使用說明)
- [技術架構](#技術架構)
- [貢獻](#貢獻)

## 簡介

LotteryService 提供了兩個主要功能：

1. **Login API**：用戶可以登錄到系統以進行身份驗證。
2. **Bet API**：用戶可以進行彩票投注，並透過區塊鏈記錄每一筆投注。

## 安裝

首先，確保你已經安裝了 Go 1.18 或更高版本。接著，按照以下步驟進行安裝：

1. 克隆此倉庫：
   ```bash
   git clone https://github.com/Jason1950/LotteryService.git
   ```
