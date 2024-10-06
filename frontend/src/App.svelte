<script lang="ts">
  import logo from './assets/images/logo-lugi-white.png'
  import {Greet} from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter the competition name below 👇"
  let name: string
  let create: string
  let search: string

  // Function to alert the value of 'Create' input
  function alertCreate() {
    alert(`Create Input: ${create}`);
  }

  // Function to alert the value of 'Search' input
  function alertSearch() {
    alert(`Search Input: ${search}`);
  }

</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="result" id="result">{resultText}</div>
<!--  <div class="input-box" id="input">-->
<!--    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>-->
<!--    <button class="btn" on:click={greet}>Greet</button>-->
<!--    <button class="btn" on:click={greet}>Greet2</button>-->
<!--  </div>-->
  <div class="input-box" id="inputcreate">
    <button class="btn" on:click={alertCreate}>Create</button>
    <input autocomplete="off" bind:value={create} class="input" id="create" type="text"/>
  </div>
  <div class="input-box" id="inputsearch">
    <button class="btn" on:click={alertSearch}>Search</button>
    <input autocomplete="off" bind:value={search} class="input" id="search" type="text"/>
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 25%;
    height: 25%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
    color: black;
  }

  .input-box .btn {
    padding: 10%;
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
