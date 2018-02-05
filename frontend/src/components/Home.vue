<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <h1 class="title is-4">
          Bienvenue sur la plateforme de gestion de vos E-Coins !
        </h1>
      </div>
    </section>

    <div class="content">
      <h3 class="home-title is-4">
        À quoi ça sert ?
      </h3>

      <hr/>

      <p>
        La finalité du E-Coin est actuellement la reconnaissance par les pairs
        (et par E-Corp).
        Ainsi, quand un collaborateur fait un geste envers vous qui vous est
        utile dans votre travail, vous pouvez lui transfèrer des E-Coins.
      </p>
      <p>
        Par exemple, Alice a débloqué Bob dans sa compréhension du dernier
        framework à la mode, du coup pour la remercier, Bob lui transfert 3 E-Coins.
      </p>
      <p>
        Autre exemple, Charles a rempli le CRA de Dorothée qui avait oublié et
        qui était en vacance. Dorothée lui transfert alors 5 E-Coins.
      </p>
      <p>
        Il n'y a pas de grille permettant de savoir combien de E-Coins donner
        pour un service précis. Ce choix est personnel. D'ailleurs, rien ne vous
        oblige à donner des E-Coins.
      </p>
    </div>

    <div class="content">
      <h3 class="home-title">
        Bleu, Rouge…?!
      </h3>

      <hr/>

      <p>
        Le E-Coin existe en plusieurs couleurs. Actuellement, il n'existe que
        Bleu et Rouge. <br>
      </p>
      <p>
        La couleur <span class="tag is-info">Bleu</span>
        correspond à la monnaie de « base » obtenue par tout le monde
        régulièrement.<br/>
      </p>
      <p>
        La couleur <span class="tag is-danger">Rouge</span> est obtenue quand
        quelqu'un vous transfert des E-Coins.
      </p>
      <p>
        Les E-Coins <span class="tag is-danger">Rouge</span> sont donc la
        représentation des services que vous avez rendus à des E-Corp.
        Ce sont les E-Coins dont vous pouvez être fiers ! Y compris en entretien
        annuel ;)
      </p>
    </div>

    <div class="content">
      <h3 class="home-title">
        Quel est mon solde ?
      </h3>

      <hr/>

      <p>
        Votre solde est indiqué en haut à droite<br/>
        Actuellement vous disposez de
        <span class="tag is-info">{{ assets.blue }}</span> E-Coins bleus,
        ainsi que de <span class="tag is-danger">{{ assets.red }}</span>
        E-Coins rouge
      </p>
    </div>

    <div class="content">
      <h3>Comment ça marche ?</h3>
      <p>
        Le E-Coin est une crypto-monnaie similaire au
        <a href="https://fr.wikipedia.org/wiki/Bitcoin">Bitcoin</a>.<br/>
        En effet, le E-Coin est basé sur une blockchain proche du Bitcoin,
        tournant grâce à <a href="https://www.multichain.com/">Multichain</a>
        <br/>
        Le livre blanc expliquant plus en détail le fonctionnement du E-Coin
        est disponible à cette
        <a href="https://github.com/e-corp/e-coin/whitepaper">adresse</a>
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
