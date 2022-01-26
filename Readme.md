# Email Unsubscriber


It is really hard to unsubscribe from Gmail because we might have many emails. The aim of this project is to easily provide way to unsubscribe email by creating your own app and run the code from opensource.
There are many tools already built but you might be concerned about the  privacy because this process requires you to give access to your inbox.

With this project everything is owned by your local machine i.e

* You will use your created google app
* You will use your own email token
* Your data will be saved in local database postgres


## Getting Google Credentials
* Create a Google Cloud Project (https://console.cloud.google.com/home/dashboard)
* Configure Oauth Consent Screen ![Oauth consent screen](https://nimb.ws/oAdiVm)
* Fill out the consent form and give use your email that you want to scan for ![](https://nimb.ws/fJkjSM) and click save and add your email as test email in next wizard
* Create Credentials for Oauth Client Id like ![](https://nimb.ws/xwLXv0)  use same `localhost`  as authorized & redirect uris
* Download oauth client json file ![](https://nimb.ws/Y2yEaS)
* Copy same json file and rename it to `credentials.json` and place it inside the project root which will be used by our application to generate `token.json`
* Enable Gmail Api Access ![](https://nimb.ws/bUsfE5)by following this link and then enable it
* Run `make publish` it will print your the link that you can use to generate token. After you see the click kindly click it
* Select your email and allow permission to app if everything went right you will be redirected to something like `https://localhost/callback?state=state-token&code=4/0AX_KGFABMBFMNBXCZXCwbxzbckjasgdiasug&scope=https://www.googleapis.com/auth/gmail.readonly`
*  copy the code `4/0AX_KGFABMBFMNBXCZXCwbxzbckjasgdiasug`  which will look like this and paste in CLI. If everything goes well it will generate a  `token.json` file which will be used on every authentication



## Usage

* Run `make publish` it will read all the email from your inbox and put those emails on RabbitMq to parse the unsubscribe link
* Run `make consume` it will consume all the queue published by first command and start parsing the link
* Run `make serve-frontend` which will present you the service name, sender email & the unsubscribe link which you can simply click and unsubscribe

## Things to do

- Write test


