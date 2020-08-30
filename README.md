# Goalist

https://run-together-towards-goals.herokuapp.com/


## 🗻Tech
- Golang
  - Gin
  - Gorm
- MySQL
- Redis
- Docker（docker-compose）
- Heroku

## 🏃How to setup
1. 

    git clone https://github.com/hariNEzuMI928/run-together-towards-goals.git goalist

2. 

    cd goalist

3. 

    docker-compose up -d

4. 

    http://localhost:8080/


## ❓About app

- 目標を構造的に定義し、その目標を達成するためのスモールSNS
- 尊厳欲求（承認欲求）・自己実現欲求を活発化させ、各々の目標を達成することをコミュニティ全体で応援する
- 自習室効果促進。「他のみんなも頑張っているから私も頑張ろう」

## 🔧機能概要

### 💻WEB
- ユーザー（users）登録・編集
- 各ユーザーは目標を立て、それに対するTODOを期限付きで立てる
- そうすることで、目標と、それを達成するためのTODO（todo_lists）を構造的に定義する
- ユーザーは日々のKPT（daily_kpts）を投稿する
- 他ユーザー情報と目標（goals）とTODO（todo_lists）を閲覧
- Goodアクション、Fightアクション

### 📪API
- ユーザー情報（user）登録・編集・取得
- 目標（.my_goals）の設定・更新
- 日々のKPT（daily_kpts）を投稿する
- Goodアクション、Fightアクション

### 📱クライアント（予定）
- 設定時間になったらKPTを投稿するよう催促
- 目標を達成したユーザーは、全ユーザーからGoodアクションが送られる
- 目標設定に問題があると感じるユーザーに対し（匿名で）コメントを送れる
