<script>
export let content = ''
export let image = ''
export let alignment = 'horizontal'
</script>

<style lang="postcss">
div > img {
    width: 320px;
}

div.horizontal > img {
    max-width: 50%;
}

div.vertical {
    height: 100vh;
    flex-direction: column;
    max-width: 75%;
    margin: 0 auto;
}

div.vertical > img {
    max-width: 75%;
}
</style>

<div class="{alignment} max-w-lg mx-auto flex items-center justify-center overflow-hidden">
    <img class="{alignment == 'horizontal' ? 'sm:pr-12' : 'pb-6'} drop-shadow-md hover:drop-shadow-lg transition-all" src={image} alt=""/>
    <p class="text-center sm:text-3xl font-extrabold text-gray-800">
        {@html content}
    </p>
</div>