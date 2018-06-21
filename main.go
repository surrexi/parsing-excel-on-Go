package main

var config TomlConfig

func main() {
    InitConfig()

    ReadFile("./URL.xlsx")
}
