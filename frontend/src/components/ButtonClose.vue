<script setup>
const close = () => {
    console.log("closing")
    window.runtime.Quit()
}
</script>

<template>
    <button 
        class="p-2 border bg-background border rounded-full hover:bg-gray-800"
        @click="close">
        <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-x w-5" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
            <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
            <path d="M18 6l-12 12" />
            <path d="M6 6l12 12" />
        </svg>
    </button>
</template>