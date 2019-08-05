<template>
    <div class="content">
        <div class="header">
            <slot name="header"></slot>
        </div>
        <div class="body">
            <slot name="body"></slot>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Content"
    }
</script>

<style scoped>
    .header {
        padding: 15px;
        text-align: right;
        border-bottom: 1px solid #ececec;
    }
    .body {
        padding: 15px;
    }
</style>
