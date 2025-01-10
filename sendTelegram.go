package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const botToken = "" // Substitua pelo seu token do bot
const chatID = ""  // Substitua pelo ID do seu chat ou grupo

// Função para enviar a mensagem para o Telegram
func sendMessage(message string) error {
	// Escapar os backticks para Markdown
	message = "```\n" + message + "\n```"  // Adiciona a formatação para o bloco de código

	// Formatar a URL da API do Telegram para envio via POST
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	// Estruturar os dados para enviar, incluindo o parâmetro parse_mode como "Markdown"
	data := map[string]interface{}{
		"chat_id":   chatID,
		"text":      message,
		"parse_mode": "Markdown",  // Definindo o formato da mensagem como Markdown
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erro ao estruturar dados para a requisição: %v", err)
	}

	// Fazer a requisição HTTP POST
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erro ao enviar mensagem: %v", err)
	}
	defer resp.Body.Close()

	// Ler a resposta (opcional)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	//    return fmt.Errorf("erro ao ler resposta: %v", err)
	// }

	// Exibir resposta da API do Telegram (opcional)
	// fmt.Println("Resposta da API:", string(body))

	return nil
}

// Função para verificar se há dados no stdin
func stdinPipe() (string, error) {
	// Verifica se há dados no stdin
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("erro ao acessar a entrada padrão: %v", err)
	}

	// Se o programa está sendo executado com dados via pipe
	if fi.Mode()&os.ModeCharDevice == 0 {
		// Cria um buffer para armazenar o conteúdo
		var buffer bytes.Buffer

		// Copia todos os dados da entrada padrão (pipe) para o buffer
		_, err := io.Copy(&buffer, os.Stdin)
		if err != nil {
			return "", fmt.Errorf("erro ao ler dados do stdin: %v", err)
		}

		return buffer.String(), nil
	}

	// Retorna uma string vazia caso não haja dados no pipe
	return "", nil
}

func main() {
	// Definindo as flags de linha de comando
	helpFlag := flag.Bool("h", false, "Exibe esta ajuda") // Flag -h
	messageFlag := flag.String("m", "", "Mensagem a ser enviada para o Telegram")

	// Processa as flags
	flag.Parse()

	// Exibe a ajuda se a flag -h for passada
	if *helpFlag {
		flag.Usage() // Exibe o usage padrão da flag
		return
	}

	// Se não houver nenhum argumento, verifica se tem entrada de pipe
	if *messageFlag != "" {
		// Caso haja um argumento -message, envia a mensagem
		err := sendMessage(*messageFlag)
		if err != nil {
			log.Fatal("Erro ao enviar mensagem:", err)
		}
	} else {
		// Caso não haja parâmetros, verifica se há dados via pipe
		if pipeMessage, err := stdinPipe(); err == nil && pipeMessage != "" {
			// Se houver dados via pipe, envia a mensagem
			err = sendMessage(pipeMessage)
			if err != nil {
				log.Fatal("Erro ao enviar mensagem do pipe:", err)
			}
		} else {
			// Caso não haja parâmetro ou pipe, exibe uma mensagem de erro
			flag.Usage() // Exibe o usage padrão da flag
		}
	}
}
