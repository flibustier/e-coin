<template>
  <div>
    <div class="wrap-banner">
      <h2>E-Coin Wallet</h2>
      <button class="btn-si" @click="login">Sign In</button>
    </div>
    <vue-particles
      color="#F4EEE7"
      :particleOpacity="0.7"
      linesColor="#F4EEE7"
      :particlesNumber="80"
      shapeType="circle"
      :particleSize="5"
      :linesWidth="2"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="3"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push">
    </vue-particles>
  </div>
</template>

<script>
  import { isLoggedIn, login } from '../auth'

  export default {
    name: 'login',
    data () {
      return {

      }
    },
    methods: {
      login () {
        login()
      }
    },
    mounted() {
      if (isLoggedIn()) {
        this.$router.push('home')
      }
    }
  }
</script>

<style lang="css">
  html, body {
    height: 100vh;
  }

  body {
    margin: 0;
    background:
    linear-gradient(
      rgba(20, 30, 48, 0.45),
      rgba(20, 30, 48, 0.45)
      ), url("../assets/img/sky.jpg");
    background-size: cover;
  }

  .wrap-banner {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
    width: 100%;
    position: absolute;
    color: #fff;
  }

  h2 {
    font-size: 7rem;
    font-family: 'Teko', sans-serif;
    text-shadow: 8px 8px #000000;
    text-transform: uppercase;
    margin: 0;
  }

  .btn-si {
    background: #dd4b39 url("../assets/img/sign-in.svg") no-repeat 1em;
    background-size: 2em;
    border-radius: 0.5em;
    border: none;
    color: white;
    fill: white;
    cursor: pointer;
    font-size: 1em;
    height: 4em;
    line-height: 1em;
    padding: 0 2em 0 4em;
    text-decoration: none;
    transition: all 0.5s;
  }

  .btn-si:hover {
    background-color: #e47365;
  }
  .btn-si:active {
    background-color: #c23321;
  }

</style>
