<template>
    <div class="sort">
        <span v-for="(item, index) in sortList"  :key="index" class="sort-item" :class="activeClass(item)" @click="toggleItem(item)">{{item.text}}</span>
    </div>
</template>

<script>
    export default {
        name: "Sort",
        data(){
            return {
                current:1,
                sortList: [{text:"Top Post", value:1}, {text:"Newest Post", value:0}]
            }
        },
        methods:{
            // eslint-disable-next-line no-unused-vars
            toggleItem(item) {
                this.current = item.value;
                this.$emit("click",  this.current)
            },
            activeClass(item) {
                const isSelected = item.value === this.current;
                return {'active':isSelected}
            }
        }
    }
</script>

<style scoped>
    .sort {
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;

        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;

        -webkit-box-pack: center;
        -ms-flex-pack: center;
        justify-content: center;

        margin-bottom: 0.8rem;
    }

    .sort-item {
        font-size: 0.8rem;
        color: #6e6e6e;
        margin: 1.5rem 0.8rem;
        padding: 0.2rem 0.8rem;
        cursor: pointer;
    }

    .sort-item.active {
        border-radius: 0.8rem;
        background-color: #0086cc;
        color: #fff;
    }

</style>