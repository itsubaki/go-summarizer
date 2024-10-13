package openai

import (
    "context"
    "fmt"
    "os"

    "github.com/openai/openai-go"
)

var defaultModel = openai.ChatModelGPT4oMini

func getModel() string {
    model := os.Getenv("OPENAI_MODEL")
    if model == "" {
        model = defaultModel
    }
    return model
}

func CallOpenAI(text string, lang string) (string, error) {
    // defaults to os.LookupEnv("OPENAI_API_KEY")
    client := openai.NewClient()
    model := getModel()
    chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
        Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
            openai.UserMessage(fmt.Sprintf("I need a summary of this text in %s:\n%s", lang, text)),
        }),
        Model: openai.F(model),
    })
    if err != nil {
        return "", err
    }

    return chatCompletion.Choices[0].Message.Content, nil
}
