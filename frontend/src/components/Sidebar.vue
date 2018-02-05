<template>
  <aside class="menu">
    <p class="menu-label">
      General
    </p>
    <ul class="menu-list">
      <li><router-link to="/" exact>Accueil</router-link></li>
    </ul>
    <p class="menu-label">
      Transactions
    </p>
    <ul class="menu-list">
      <li><router-link to="/history">Historique</router-link></li>
      <li><router-link to="/transfer">Transfert</router-link></li>
    </ul>
  </aside>
</template>
