<p align="center">
    <a href="https://e-coin.jonathan.pl/" target="_blank">
        <img width="100" src="./resources/ecoin-dark.png" alt="E-Coin">
    </a>
</p>

---

[multichain]: https://github.com/MultiChain/multichain

#### What is E-Coin?

E-Coin is a centralised and corporate digital assets for your E-Corp employees.

#### Centralised? What about blockchain?

E-Coin use blockchain as a storage for storing assets and transactions.

The blockchain is powered by [Multichain]
for now, perfect for building a private blockchain!

#### How is it working?

E-Coin platform use a [Vue.js](https://github.com/vuejs/vue) front-end, and a Go
back-end (WIP).

![E-Coin architecture](./resources/e-coin-archi.png)

#### Can I test it?

Sure! As a loyal E-Corp employee, you can register and get your first E-Coin
immediately!

Go to: <https://e-coin.jonathan.pl> to claim your reward!

#### It's awesome! I want to build my own E-Coin!

Sure! 

##### Build Blockchain

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

##### Build Back-end

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

##### Build Front-end

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
  - [ ] Get list of user
  - [x] Create new transaction
  - [ ] Get user history 
- [ ] Fully functional
- [x] Release
- [ ] Add a blockchain explorer
- [ ] Add new functionality

#### License

[MIT](http://opensource.org/licenses/MIT)
