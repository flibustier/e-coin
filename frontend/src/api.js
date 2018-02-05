import axios from "axios";
import config from "./config.json";

// Get user's assets balance
const balance = () => Promise.resolve({
	blue: 10,
	red: 10
});
/*	axios
		.get(`${config.api}/user/balance`)
		.then(response => {
			const assets = {
				blue: 0,
				red: 0,
			};

			response.data.assets.forEach(asset => {
				if (asset.name !== undefined) {
					assets[asset.name.toLowerCase()] = asset.qty;
				}
			});
			return assets;
		})
		.catch(e => {
			throw e;
		});
*/
// Get users list
const users = () =>
	axios
		.get(`${config.api}/users`)
		.then(response => {
			return response.data.map(user => user.email);
		})
		.catch(e => {
			throw e;
		});

// -- Transfer

const transfer = (beneficiary, assets) =>
	axios
		.post(`${config.api}/user/transfer`, {
			emailTo: beneficiary,
			balance: {
				assets: assets.map(asset => {
					return {
						name: asset.currency.toUpperCase(),
						qty: asset.amount,
					};
				}),
			},
		})
		.catch(e => {
			throw e;
		});

// -- Account

const importAddress = (address, privkey) =>
	axios
		.post(`${config.api}/user/mining/import`, {
			publicKey: address,
			privateKey: privkey,
		})
		.catch(e => {
			throw e;
		});

const register = address =>
	axios
		.post(`${config.api}/user/mining/initialize`, {
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
	axios.post(`${config.api}/user/goodies/buy`, goodie).catch(e => {
		throw e;
	});

// -- History

const transactions = () =>
	axios
		.get(`${config.api}/user/transactions`)
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
	axios.get(`${config.api}/user/exchange?number=${amount}`).catch(e => {
		throw e;
	});

export default {
	balance,
	users,
	transfer,
	importAddress,
	register,
	goodies,
	purchase,
	transactions,
	exchangeRate,
	change,
};
