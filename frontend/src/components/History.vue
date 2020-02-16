<template>
  <div>
    <div class="is-centered has-text-centered" v-if="isFetching">
      <img src="../assets/img/loading.svg">
    </div>

    <div v-if="!isFetching && transactions.length < 1">
      <span>No data found for this user</span>
    </div>

    <section v-if="!isFetching && transactions.length > 0">
      <b-table
        :data="transactions"
        paginated:true
        per-page="10"
        default-sort-direction="desc"
        default-sort="time">

        <template slot-scope="props">

          <b-table-column field="from" label="Origin" sortable>
            <b-tag :type="getType(props.row.from)" rounded>
              {{ props.row.from }}
            </b-tag>
          </b-table-column>

          <b-table-column field="to" label="Recipient" sortable>
            <b-tag :type="getType(props.row.to)" rounded>
              {{ props.row.to }}
            </b-tag>
          </b-table-column>

          <b-table-column field="amount" label="Assets" sortable>
            <div style="display: flex; flex-direction: row" v-for="asset in props.row.assets" :key="asset.name">
              {{ Math.abs(asset.qty) }}
              <img v-if="asset.name === 'blue'" class="currency" src="../assets/img/ecoin-blue.png"/>
              <img v-else class="currency" src="../assets/img/ecoin-red.png"/>
            </div>
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
        isFetching: true,
        address: null
      };
    },

    methods: {
      getType(a) {
        return (a === this.address) ? "is-info" : "is-white";
      }
    },

    created() {
      api.transactions().then(transactions => {
        this.transactions = transactions;
        this.isFetching = false;
      }).catch(e => {
        this.isFetching = false;
        this.$toast.open({
          duration: 5000,
          message: `Oops! ${e}`,
          position: 'is-bottom',
          type: 'is-danger'
        })
      });
      api.address().then(a => this.address = a);
    },
  };
</script>

<style scoped>
  .currency {
    height: 1.5em;
  }
</style>
