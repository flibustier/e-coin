import axios from "axios";
import config from "./config.json";
import { getAccessToken, getIdToken } from "./auth";

axios.defaults.headers.common["Authorization"] = "Bearer " + getAccessToken();
axios.defaults.headers.common["id_token"] = getIdToken();

// Get user's assets balance
const balance = () =>
  axios
    .get(`${config.api}/users/balance`)
    .then(response => {
      const assets = {
        blue: 0,
        red: 0
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
          qty: asset.amount
        };
      })
    })
    .catch(e => {
      throw e;
    });

// -- Account

const importAddress = (address, privkey) =>
  axios
    .post(`${config.api}/users/mining/import`, {
      publicKey: address,
      privateKey: privkey
    })
    .catch(e => {
      throw e;
    });

const register = address =>
  axios
    .post(`${config.api}/users/mining/initialize`, {
      publicKey: address,
      privateKey: null
    })
    .catch(e => {
      throw e;
    });

// -- History

const transactions = () =>
  axios
    .get(`${config.api}/users/transactions`)
    .then(response => {
      return response.data
        .filter(t => t.balance.assets.length > 0)
        .map(transaction => {
          let from = transaction.myaddresses[0];
          let to = transaction.addresses[0];
          if (transaction.balance.assets[0].qty > 0) {
            to = from;
            from = transaction.addresses[0];
          }

          return {
            from: from,
            to: to,
            time: transaction.time * 1000,
            assets: transaction.balance.assets
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
  importAddress,
  purchase,
  register,
  transactions,
  transfer,
  users
};
