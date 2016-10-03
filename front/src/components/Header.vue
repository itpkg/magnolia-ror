<template>
<nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
  <a class="navbar-brand" href="#">Project name</a>
  <ul class="nav navbar-nav">
    <li class="nav-item active">
      <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="#">About</a>
    </li>
    <li class="nav-item">
      <a class="nav-link" href="#">Contact</a>
    </li>
    <li class="nav-item">
      <router-link to="/users/sign-in">Sign in</router-link>
    </li>
  </ul>
</nav>
</template>

<script>
export default {
  data () {
    return {
      msg: 'Header'
    }
  }
}
</script>
