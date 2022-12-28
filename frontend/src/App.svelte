<script>
  import svelteLogo from './assets/svelte.svg'
  import Counter from './lib/components/Counter.svelte'
  import "carbon-components-svelte/css/g90.css";
  import { Tabs, Tab, TabContent } from "carbon-components-svelte";
  import Navbar from './lib/layout/Navbar.svelte';
  import Products from './lib/pages/Products.svelte';
  import Orders from './lib/pages/Orders.svelte';
  import About from './lib/pages/About.svelte';
</script>



<main>
  <Navbar></Navbar>
  <div class="content">
    <Tabs>
      <Tab label="Products" />
      <Tab label="Orders" />
      <Tab label="About" />
      <svelte:fragment slot="content">
        <TabContent>
          <Products></Products>
        </TabContent>
        <TabContent>
          <Orders></Orders>
        </TabContent>
        <TabContent>
          <About></About>
        </TabContent>
      </svelte:fragment>
    </Tabs>
  </div>
  
</main>

<style>
  .content {
    display: block;
    position: fixed;
    width: 100%;
    height: calc(100% - var(--navbar-height));
    top: var(--navbar-height);
    left: 0;
    overflow-y: scroll;
    z-index: 9999;
    padding: 2rem;
  }

</style>
