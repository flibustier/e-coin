<template>
  <nav class="navbar has-shadow">
    <div class="navbar-brand">
      <router-link class="navbar-item" id="brand-title" to="/">
        E-Coin
      </router-link>
    </div>
    <div class="navbar-end">
      <div class="navbar-item">
        <p class="is-primary">Bleu : {{ assets.blue }}</p>
        <b-icon
                icon="euro_symbol"
                size="is-large"
                type="is-info">
        </b-icon>
      </div>
      <div class="navbar-item">
        <p class="is-primary">Rouge : {{ assets.red }}</p>
        <b-icon
                icon="euro_symbol"
                size="is-large"
                type="is-danger">
        </b-icon>
      </div>
      <span class="navbar-item">
        <i class="fa fa-sign-out" @click="logout()"></i>
      </span>
    </div>
  </nav>
</template>

<script>
    import { logout } from '../auth'

    export default {
        props: ["assets"],
        name: 'navbar',
        methods: {
            logout () {
                logout();
                this.$router.go('/login')
            }
        },
    }
</script>

<style>

  #brand-title {
    font-size: 3rem;
    font-family: 'Teko', sans-serif;
    color: #B31835;
    text-shadow: 3px 3px #000;
    text-transform: uppercase;
    margin: 0;
  }

</style>
