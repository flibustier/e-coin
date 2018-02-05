<template>
  <div>
    <h1 class="title is-4">
      Transfert de E-Coins
    </h1>

    <hr/>

    <b-notification type="is-warning">
      Tout transfert est immédiat et <strong>irréversible</strong>.
    </b-notification>

    <h2 class="subtitle">
      Destinataire
    </h2>

    <b-field grouped>
      <b-field :type="validateBeneficiary">
        <b-autocomplete
        v-model="autocomplete.name"
        :data="filteredDataArray"
        field="email"
        placeholder="Ex: Jon…"
        icon="search"
        type="search"
        :loading="autocomplete.isFetching"
        @select="option => autocomplete.selected = option">
      </b-autocomplete>
    </b-field>
  </b-field>

  <hr/>

  <h3 class="subtitle">Montants</h3>

  <div class="tile is-ancestor">
    <div v-for="asset in wallet" class="tile is-parent is-4">
      <article class="tile is-child box">
        <h1 class="subtitle">E-Coin {{ asset.name }}</h1>
        <div class="block">
          <p>
            <slider :type="asset.color" :value="asset.amount" :max="asset.max"
            :step="1" is-fullwidth @change="update($event, asset)"></slider>
          </p>
          <p>
            <b-input type="number" class="is-small" v-model="asset.amount"
            min="0" :max="asset.max"></b-input>
          </p>
        </div>
      </article>
    </div>
  </div>

  <hr/>

  <b-field>
    <p class="control level-right">
      <button class="button is-warning" @click="isModalActive = true"
      :disabled="!ready">
        <span>Transfert</span>
      </button>
    </p>
  </b-field>

  <b-modal :active.sync="isModalActive" :width="640">
    <form>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Confirmation de transfert</p>
        </header>
        <section class="modal-card-body">
          <p>
            Vous êtes sur le point d'effectuer le transfert suivant à
            {{ autocomplete.selected }}
          </p>
          <p v-for="asset in wallet" v-if="asset.amount > 0">
            {{ asset.amount }} {{ asset.name}}
          </p>
        </section>
        <footer class="modal-card-foot">
          <button class="button is-danger" type="button"
          @click="isModalActive = false">Annuler</button>
          <button class="button is-primary" type="button"
          @click="validate">Confirmer</button>
        </footer>
      </div>
    </form>
  </b-modal>
</div>
</template>

<script>
import api from "../api";
import Slider from "vue-bulma-slider";

export default {
  components: {
    Slider,
  },

  props: ["assets"],

  data() {
    return {
      autocomplete: {
        data: [],
        name: "",
        selected: null,
        isFetching: false,
      },
      isModalActive: false,
      isLoading: false,
      wallet: [
        {
          name: "Bleu",
          currency: "blue",
          amount: 0,
          max: 0,
          color: "",
        },
        {
          name: "Rouge",
          currency: "red",
          amount: 0,
          max: 0,
          color: "danger",
        },
      ],
    };
  },

  watch: {
    assets: function() {
      this.updateWallet();
    },
  },

  methods: {
    validate() {
      this.isModalActive = false;
      const loadingComponent = this.$loading.open();

      api
      .transfer(this.autocomplete.selected, this.wallet)
      .then(() => {
        loadingComponent.close();
        this.$emit("transaction");
        this.$toast.open({
          duration: 5000,
          message: `Votre transfert s'est déroulé avec succès!`,
          type: "is-success",
        });
      })
      .catch(() => {
        loadingComponent.close();
        this.$toast.open({
          duration: 5000,
          message: `Transfert échoué, merci de réessayer plus tard…`,
          type: "is-danger",
        });
      });
    },

    update(val, asset) {
      asset.amount = Number(val);
    },

    updateWallet() {
      this.wallet.forEach(asset => {
        asset.max = this.assets[asset.currency];
      });
    },
  },

  computed: {
    filteredDataArray() {
      return this.autocomplete.data.filter(option => {
        return (
          option
          .toString()
          .toLowerCase()
          .indexOf(this.autocomplete.name.toLowerCase()) >= 0
        );
      });
    },
    ready() {
      return (
        this.autocomplete.selected &&
        this.wallet.filter(asset => {
          return asset.amount > 0;
        }).length > 0
      );
    },
    validateBeneficiary() {
      return this.autocomplete.selected ? "is-success" : "is-info";
    },
  },
  created() {
    this.autocomplete.data = [];
    this.autocomplete.isFetching = true;
    api.users().then(users => {
      this.autocomplete.data = users;
      this.autocomplete.isFetching = false;
    });
    this.updateWallet();
  },
};
</script>

<style lang="scss" scoped>
p {
  margin-bottom: 20px;
}
</style>
