<p align="center">
    <a href="https://e-coin.jonathan.pl/" target="_blank">
        <img width="100" src="./resources/ecoin-dark.png" alt="E-Coin">
    </a>
</p>

---

[multichain]: https://github.com/MultiChain/multichain

#### What is E-Coin?

E-Coin is a digital asset for your E-Corp employees. It's also a demo 
application.

#### What about blockchain?

E-Coin use blockchain as a storage for storing assets and transactions.

The blockchain is powered by [Multichain]
for now, perfect for building a private blockchain!

#### How is it working?

E-Coin platform use a [Vue.js](https://github.com/vuejs/vue) front-end, a Go
back-end interacting with Multichain, and Auth0 for managing authentication.

<p align="center">
    <img alt="E-Coin architecture" src="./resources/e-coin-archi.png">
</p>

#### Can I test it?

Sure! As a loyal E-Corp employee, you can register and get your first E-Coin
immediately!

Go to: <https://e-coin.jonathan.pl> to claim your reward!

> You can register with any email address in order to get a _e-coin wallet address_ and send some assets 
to another address, but if you don't want to register your email, you can test the 
application using `demo@jonathan.pl` email and `demo` password! 

You can also explore the Blockchain with the [Explorer](https://explorer.jonathan.pl)!

#### Awesome! I want to develop my own E-Coin!

Sure! You will need [Golang](https://golang.org/dl/) and [NPM](https://nodejs.org/en/)

##### Automated

If you are on Windows, just launch `./start.bat`

##### Manually

###### Build your Blockchain

First you need to download [Multichain], 
then follow these steps by replacing `e-coin` with the name of your currency:
```$xslt
multichain-util create e-coin

[...]
You can edit it in /home/t.wellick/.multichain/e-coin/params.dat before running multichaind for the first time.
[...]
```

You can adjust the parameters in `params.dat` file, you can take inspiration from [those](./resources/params.dat)

Next, it's time to run your blockchain node
```
multichaind e-coin

MultiChain 2.0 alpha 2 Daemon (latest protocol 20002)

Other nodes can connect to this node using:
multichaind e-coin@42.42.42.42:6667

This host has multiple IP addresses, so from some networks:
multichaind e-coin@42.42.42.42:6667

Listening for API requests on port 6666 (local only - see rpcallowip setting)

Node ready.
```

You can now fill you `.env` file with the RPC credentials in
`/home/t.wellick/.multichain/e-coin/multichain.conf` file

###### Build Back-end

```$xslt
go get github.com/flibustier/e-coin
```

Don't forget to fill the `.env` file

The dependencies are:
* gorilla/mux
* gorilla/context
* rs/cors
* joho/godotenv
* boltdb/bolt
* golangdaddy/multichain-client
* auth0-community/go-auth0
* gopkg.in/square/go-jose.v2

###### Build Front-end

Follow instructions in the [frontend directory](./frontend/README.md)
```
npm install
npm run build
```

---

#### Roadmap

- [x] Static front-end
- [x] Go back-end
  - [x] Multichain integration
  - [x] Database support
  - [x] Create new users
  - [x] Get user balance
  - [x] Get list of user
  - [x] Create new transaction
  - [x] Get user history 
- [x] Fully functional
- [x] Release
- [x] Add a blockchain explorer
- [ ] Add new functionality

#### License

[MIT](http://opensource.org/licenses/MIT)
