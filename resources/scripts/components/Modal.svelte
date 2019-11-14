<script>
export let fadeIn = true
</script>

<style lang="postcss">
.overlay {
    background-image: url("data:image/svg+xml,%3Csvg width='12' height='16' viewBox='0 0 12 16' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M4 .99C4 .445 4.444 0 5 0c.552 0 1 .45 1 .99v4.02C6 5.555 5.556 6 5 6c-.552 0-1-.45-1-.99V.99zm6 8c0-.546.444-.99 1-.99.552 0 1 .45 1 .99v4.02c0 .546-.444.99-1 .99-.552 0-1-.45-1-.99V8.99z' fill='%239C92AC' fill-opacity='0.1' fill-rule='evenodd'/%3E%3C/svg%3E");
    animation-duration: 0.5s;
}

.modal {
    width: 90%;
    max-width: 400px;
    min-width: 320px;
}
</style>

<div class="fixed top-0 left-0 w-full h-full z-50 flex items-center justify-center overflow-y-auto">
    <div class="overlay {fadeIn && 'animated fadeIn'} w-full h-full absolute bg-gray-800 z-40"></div>
    <div class="modal animated bounceIn p-4 bg-white rounded-lg shadow-md z-50">
        <slot></slot>
    </div>
</div>