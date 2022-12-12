<template>
  <div id="app">
    <Header></Header>
    <keep-alive>
      <router-view>
        <div class="nav">
          <Home></Home>
        </div>
      </router-view>
    </keep-alive>
    <side class="side" v-show="$route.meta.Show"></side>
    <Footer v-show="$route.meta.Show"></Footer>
  </div>

</template>

<script>
import Home from './views/Home'
import Header from './components/Header'
import Footer from './components/Footer'
import side from './components/side'
export default {
  name: 'App',
  components: {
    Home,
    Header,
    Footer,
    side
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  position: relative;
  width: 100vw;
  height: 100vh;
  overflow-x: hidden;
}

.nav {
  width: 100vw;
  /* margin-left: auto;
  margin-right: auto; */

}

.side {
  /* position: relative; */

  top:25vh;
  z-index: 999;
}
</style>
