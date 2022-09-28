package handlers

import (
	"fmt"
)


func MessagesHandler(msg string, user string) string {

	switch msg {
		case "Welcome":
			return welcome(user)
		case "Bye":
			return bye(user)
		default:
	}

	return ""

}

func welcome(user string) string {

	welcome_msg := fmt.Sprintf(`
Olá! %s 👋
Eu sou iSam o seu Bot 🤖
Em que posso te ajudar?
Selecione uma das opções abaixo.
`, user)

	return welcome_msg
}

func bye(user string) string {

	bye_msg := fmt.Sprintf(`
 %s, obrigado por utilizar nossos serviços.

 Se precisar de mais algum auxílio, clico em uma das opção abaixo.

 iSamBot agradece a preferencia.
`, user)

	return bye_msg
}

// var TextMsg = map[string]string{
//
// 	"welcome": fmt.Sprintf(`
// Olá! 👋 ...
// Eu sou iSam o seu Bot 🤖
// Em que posso te ajudar?
// Selecione uma das opções abaixo.
// `),
//
// 	"help": `
// Olá,
// Percebo que precisa de ajuda, em que posso ajudar?
//
// `,
//
// 	"bye": `
//  Obrigado por utilizar nossos serviços.
//  Se precisar de mais algum auxílio, clico em uma das opção abaixo.
//
//  iSamBot agradece a preferencia.
// `,
//
// 	"notallowed": `
// Atenção!!!
//
// Mensagens fora do contexto do não serão permitidas.
// Somente as opções abaixo ou quando for solicitado.
//
// iSamBot, agradeve sua compreenção.
// `,
// }
