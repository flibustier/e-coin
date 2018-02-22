import axios from "axios";
import config from "./config.json";
import {getAccessToken, getIdToken} from "./auth";

axios.defaults.headers.common['Authorization'] = 'Bearer ' + getAccessToken();
axios.defaults.headers.common['id_token'] = getIdToken();

// Get user's assets balance
const balance = () =>
	axios
		.get(`${config.api}/users/balance`)
		.then(response => {
			const assets = {
				blue: 0,
				red: 0,
			};

			response.data.forEach(asset => {
				if (asset.name !== undefined) {
					assets[asset.name.toLowerCase()] = asset.qty;
				}
			});
			return assets;
		})
		.catch(e => {
			throw e;
		});

// Get users list
const users = () =>
	axios
		.get(`${config.api}/users`)
		.then(response => {
			return response.data;
		})
		.catch(e => {
			throw e;
		});

// Get user address
const address = () =>
  axios
    .get(`${config.api}/users/address`)
    .then(response => {
      return response.data;
    })
    .catch(e => {
      throw e;
    });

// -- Transfer

const transfer = (beneficiary, assets) =>
	axios
		.post(`${config.api}/users/transfer`, {
			to: beneficiary,
			assets: assets.map(asset => {
				return {
					name: asset.currency,
					qty: asset.amount,
				};
			}),
		})
		.catch(e => {
			throw e;
		});

// -- Account

const importAddress = (address, privkey) =>
	axios
		.post(`${config.api}/users/mining/import`, {
			publicKey: address,
			privateKey: privkey,
		})
		.catch(e => {
			throw e;
		});

const register = address =>
	axios
		.post(`${config.api}/users/mining/initialize`, {
			publicKey: address,
			privateKey: null,
		})
		.catch(e => {
			throw e;
		});

// -- Goodies

const goodies = () =>
	axios
		.get(`${config.api}/goodies`)
		.then(response => {
			return response.data;
		})
		.catch(e => {
			throw e;
		});

const purchase = goodie =>
	axios.post(`${config.api}/users/goodies/buy`, goodie).catch(e => {
		throw e;
	});

// -- History

const transactions = () =>
	axios
		.get(`${config.api}/users/transactions`)
		.then(response => {
			return response.data.map(transaction => {
				const unit = Object.keys(transaction.deposits)[0];
				return {
					type: transaction.type.toLowerCase(),
					from: transaction.from,
					to: transaction.to,
					time: transaction.transactionTime,
					unit: unit,
					amount: transaction.deposits[unit],
				};
			});
		})
		.catch(e => {
			throw e;
		});

// -- Exchange

const exchangeRate = () =>
	axios
		.get(`${config.api}/admin/parameters`)
		.then(response => {
			return response.data["exchangeRate"];
		})
		.catch(e => {
			throw e;
		});

const change = amount =>
	axios.get(`${config.api}/users/exchange?number=${amount}`).catch(e => {
		throw e;
	});

export default {
	address,
	balance,
  change,
  exchangeRate,
  goodies,
  importAddress,
  purchase,
  register,
  transactions,
	transfer,
  users,
};
