<template>
  <div id="app">
    <HeaderW/>
    <div class="container">      
      <FooterW/>
    </div>
  </div>
</template>

<script>
import HeaderW from './components/Header'
import FooterW from './components/Footer'
import Hello from './components/Hello'

export default {
  components: {
    Hello,
    HeaderW,
    FooterW
  }
}
</script>

<style>
body {
  padding-top: 5rem;
}
</style>
