# Golang API Sample

使用 **golang migrate** 與 **gorm** 的功能，可以快速進行資料庫的操作，搭配 **gin** 的框架，可以提升使用的效能，開發過程使用 `air` 的套件，達成在 local 或是 docker 上皆可以使用 live reload 的功能

---

### 使用的套件
1. [golang](https://github.com/golang)/[go](https://github.com/golang/go)
2. [golang-migrate](https://github.com/golang-migrate)/[migrate](https://github.com/golang-migrate/migrate)
3. [go-gorm](https://github.com/go-gorm)/[gorm](https://github.com/go-gorm/gorm)
4. [cosmtrek](https://github.com/cosmtrek)/[air](https://github.com/cosmtrek/air)
5. [gin-gonic](https://github.com/gin-gonic)/[gin](https://github.com/gin-gonic/gin)

---

## Install

- ### Golang Migrate

    > 參考資料
    > [golang-migrate install](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)

    - #### Linux

        > 參考資料
        > [Unable to install golang migrate library on Ubuntu 20.4](https://stackoverflow.com/questions/66621682/unable-to-install-golang-migrate-library-on-ubuntu-20-4/69478562#69478562)
        > [Golang Migrate Release](https://github.com/golang-migrate/migrate/releases)

        ```bash
        # 選擇要安裝的版本與環境
        # curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
        $ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

        
        $ mv migrate.linux-amd64 $GOPATH/bin/migrate
        ```

    - #### Mac

        ```bash
        $ brew install golang-migrate
        ```

    - #### Windows

        1. 要先安裝 `scoop` 套件管理工具，參考資訊 [Windows 套件管理工具 - Scoop](https://www.gss.com.tw/blog/windows-%E5%A5%97%E4%BB%B6%E7%AE%A1%E7%90%86%E5%B7%A5%E5%85%B7-scoop)


            #### ❗ 要在 PowerShell 裡執行
            ```powershell
            > Set-ExecutionPolicy RemoteSigned -scope CurrentUser
            > iwr -useb get.scoop.sh | iex 
            ```

        2. 安裝 golang-migrate

            ```powershell
            # powershell
            > scoop install migrate
            ```

- ### Golang Gorm

    此套件類似像 `Laravel Eloquent` 的功能，可以方便進行資料庫的操作，以及建立 `Model` 之間的關聯
    <br/>

    ```bash
    $ go get -u gorm.io/gorm
    ```

- ### Golang Air

    安裝此套件可以做到 live reload 的功能，意思是，不用每次修改完檔案，都要執行 `go build`，提升開發效率 
    <br/>

    ```bash
    # binary will be $(go env GOPATH)/bin/air
    $ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

    # or install it into ./bin/
    $ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

    $ air -v # 確認是否安裝成功
    ```

    ##### Options

    建立 Air 的設定檔(`.air.toml`)，如果不執行，就會用預設的設定
    <br/>

    ```bash
    $ air init
    ```

- ### Golang Gin

    Golang 的框架，在 API 開發上效率最快
    <br/>

    ```bash
    $ go get -u github.com/gin-gonic/gin
    ```

- ### env

    先將 `.env.example` 複製以下三份檔案

    | File               | Description         |
    |:------------------ |:------------------- |
    | `.env`             | local 會使用        | 
    | `docker_app.env`   | docker app 會使用   |
    | `docker_mysql.env` | docker mysql 會使用 |


---

## Usage

- ### Local

    `air` 套件安裝完之後即可在專案目錄執行
    <br/>
	
	```bash
	$ air
	  __    _   ___  
	 / /\  | | | |_) 
	/_/--\ |_| |_| \_ , built with Go 

	watching .
	watching bin
	!exclude tmp
	building...
	running.
    ```

    ##### ❗️ Mac 無法使用 `air` 指令時，請使用以下方法

    1. 開啟 `~/.zshrc`

        ```bash
        $ vim ~/.zshrc
        ```

    2. 新增指令

        ```bash
        # ~/.zshrc
        alias air="$GOPATH/bin/air"
        ```
        
    3. 重新讀取檔案

        ```bash
        $ source ~/.zshrc
        ```

    4. 在 terminal 執行

        ```bash
        $ air
        ```

- ### Docker

    ```bash
    $ make docker-run
    ```

---

## Migrate

- ### 建立 migrate 檔案

    通常都是在 local 建立，到 docker 上只需要執行 migrate 即可
    <br/>

    ```bash
    # migrate create -ext sql -dir <migration file folder> -seq <file name>
    $ migrate create -ext sql -dir ./db/migrate -seq create_users_table
    ```

    執行完畢之後，會在指定的資料夾中自動產生 up & down 的 `.sql` 檔案，在裡面撰寫 MySQL 的語法建立 table

- ### 執行 migrate

    ❗️ 特別注意：如果是要在 docker container 中執行，請務必將 `Makefile` 中最上方的程式碼註解，否則讀取到的環境變數會是 `.env`，而非 docker 的環境變數
    <br/>

    ```Makefile
    # 如果在 docker container 中要使用的話，下面這段要註解，才可以取得 docker env variable
    # ifneq (,$(wildcard ./.env))
    #     include .env
    #     export
    # endif
    ```

    如果是要在 docker container 操作的話，要先進去 container 裡面
    ```bash
    $ docker exec -it app sh # 進入 docker container
    ```

    - #### migrate up

        ==執行之前，請先確認 database 是否已經建立==
        <br/>

        ```bash
        $ make migrate-up
        ```

    - #### migrate down

        ```bash
        $ make migrate-down
        ```

---

## API Sample

- ### Create user

    #### Request

    ```curl
    POST {base_url}/api/user/create
    ```

    ```json
    {
        "first_name": "Ming",
        "last_name": "Wang",
        "email": "ming@example.com"
    }
    ```

    #### Response

    ```json
    {
        "message": "success"
    }
    ```

- ### Retrieve all users

    #### Request

    ```curl
    GET {base_url}/api/users
    ```

    #### Response

    ```json
    [
        {
            "id": 1,
            "first_name": "Ming",
            "last_name": "Wang",
            "email": "ming@example.com",
            "created_at": "2022-02-03T14:59:21+08:00",
            "updated_at": "2022-02-03T14:59:21+08:00"
        },
        {
            "id": 2,
            "first_name": "Peter",
            "last_name": "Lee",
            "email": "peter@example.com",
            "created_at": "2022-02-03T15:03:01+08:00",
            "updated_at": "2022-02-03T15:03:01+08:00"
        },
    ]
    ```

- ### Retrieve user by id

    #### Request

    ```curl
    GET {base_url}/api/user/{id}
    ```

    #### Response

    ```json
    [
        {
            "id": 1,
            "first_name": "Ming",
            "last_name": "Wang",
            "email": "ming@example.com",
            "created_at": "2022-02-03T14:59:21+08:00",
            "updated_at": "2022-02-03T14:59:21+08:00"
        }
    ]
    ```

---