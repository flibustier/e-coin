<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <h1 class="title is-4">
          Welcome to the management platform of your E-Coins!
        </h1>
      </div>
    </section>

    <div class="content">
      <h3 class="home-title is-4">
        What's the point?
      </h3>

      <hr/>

      <p>
        The purpose of the E-Coin is currently peer recognition (and by E-Corp).
        Thus, when a collaborator makes a gesture to you that is useful to you in your work,
        you can transfer him E-Coins.
      </p>
      <p>
        For example, Alice unlocked Bob in his understanding of the latest fashionable framework,
        so to thanks her, Bob transferred her 3 E-Coins.
      </p>
      <p>
        There is no grid to know how many E-Coins give for a specific service.
        This choice is personal. Besides, nothing forces you to give E-Coins.
      </p>
    </div>

    <div class="content">
      <h3 class="home-title">
        Blue, Red…?!
      </h3>

      <hr/>

      <p>
        E-Coin exists in different colors. Currently, only blue and red are available.
      </p>
      <p>
        The <span class="tag is-info">Blue</span> color corresponds to the normal reward given
        to all E-Corp employees.
      </p>
      <p>
        The <span class="tag is-danger">Red</span> color can be obtain when someone gives you some E-Coins.
      </p>
      <p>
        The <span class="tag is-danger">Red</span> are the proof of services you gived to
        E-Corp.
        These are the E-Coins you can be proud of! Including in HR interview ;)
      </p>
    </div>

    <div class="content">
      <h3 class="home-title">
        What is my balance?
      </h3>

      <hr/>

      <p>
        Your balance can be viewed in the top right of this application<br/>
        Currently, you have
        <span class="tag is-info">{{ assets.blue }}</span> blue E-Coins,
        and <span class="tag is-danger">{{ assets.red }}</span> red E-Coins
      </p>
    </div>

    <div class="content">
      <h3>How it works?</h3>
      <p>
        E-Coins are assets on a blockchain similar to
        <a href="https://fr.wikipedia.org/wiki/Bitcoin">Bitcoin</a>,
        built with <a href="https://www.multichain.com/">Multichain</a>
        <br/>
        More information about this project can be found in
        <a href="https://github.com/flibustier/e-coin">this repository</a>
      </p>
    </div>

  </div>
</template>

<script>
export default {
  props: ["assets"],

}
</script>

<style>
.home-title {
  margin-top: 2em;
}
</style>
