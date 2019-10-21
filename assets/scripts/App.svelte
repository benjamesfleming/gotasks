<script>
    import Router from 'svelte-spa-router'
    import Navigation from '~/components/Navigation'
    import routes from '~/utils/routes'
</script>

<Navigation />
<Router {routes}/>