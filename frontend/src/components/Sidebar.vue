<template>
  <aside class="menu">
    <p class="menu-label">
      General
    </p>
    <ul class="menu-list">
      <li><router-link to="/" exact>Home</router-link></li>
    </ul>
    <p class="menu-label">
      Transactions
    </p>
    <ul class="menu-list">
      <li><router-link to="/history">History</router-link></li>
      <li><router-link to="/transfer">Transfer</router-link></li>
    </ul>
  </aside>
</template>
