<template>
  <div>
    <div class="is-centered has-text-centered" v-if="isFetching">
      <img src="../assets/img/loading.svg">
    </div>

    <div v-if="!isFetching && transactions.length < 1">
      <span>Aucune donnée trouvée pour cet utilisateur</span>
    </div>

    <section v-if="!isFetching && transactions.length > 0">
      <b-table
        :data="transactions"
        paginated:true
        per-page="10"
        default-sort-direction="desc"
        default-sort="time">

        <template slot-scope="props">

          <b-table-column label="Type de transaction" sortable>
            <div v-if="props.row.type === 'user'">
              <span class="tag is-info">
                <b-icon :icon="'swap_horiz'">Utilisateur</b-icon>
                Utilisateur
              </span>
            </div>
            <div v-else>
              <span class="tag is-danger">
                <b-icon :icon="'shopping_cart'">Utilisateur</b-icon>
                Achat
              </span>
            </div>
          </b-table-column>

          <b-table-column field="from" label="Émetteur" sortable>
            {{ props.row.from }}
          </b-table-column>

          <b-table-column field="to" label="Bénéficiaire" sortable>
            {{ props.row.to }}
          </b-table-column>

          <b-table-column field="amount" label="Montant" width="40" sortable numeric>
            {{ props.row.amount }}
          </b-table-column>

          <b-table-column field="unit" label="Devise" sortable>
            <span class="tag is-danger" v-if="props.row.unit === 'RED'">
              Rouge
            </span>
            <span class="tag is-info" v-else>
              Bleu
            </span>
          </b-table-column>

          <b-table-column field="time" label="Date" sortable centered>
            {{ new Date(props.row.time).toLocaleDateString() }}
          </b-table-column>

        </template>

      </b-table>
    </section>
  </div>
</template>

<script>
  import api from "../api";

  export default {
    data() {
      return {
        transactions: [],
        isFetching: false,
      };
    },

    created() {
      api.transactions().then(transactions => {
        this.transactions = transactions;
        this.isFetching = false;
      });
    },
  };
</script>
