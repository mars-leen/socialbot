<template>
    <div class="mut-tags">
        <div class="scroll">
            <div class="tag-box">
                <span class="tag" :class="tagClass(tag.value)" v-for="(tag, index) in tags"  :key="index" @click="toggleTag(tag.value)">
                    {{ tag.text }}
                </span>
            </div>
        </div>
    </div>
</template>

<script>
    import {getAllTagApi} from '../../api/app'
    import {Tag, Icon} from 'vant';
    import {mapGetters, mapMutations} from 'vuex';

    export default {
        name: "Tags",
        components: {
            [Tag.name]:Tag,
            [Icon.name]:Icon,
        },
        data() {
            return {
                tags:[],
            }
        },
        created(){
            this.getTags();
        },
        methods: {
            ...mapGetters({
                getSelectedTags: "getSelectedTags"
            }),
            ...mapMutations({
                updateSelectedTags : "updateSelectedTags"
            }),
            toggleTag(tagId){
                this.updateSelectedTags(tagId);
                this.$emit("click")
            },
            tagClass(tagId) {
                const isSelected = this.getSelectedTags().includes(tagId);
                return {'active':isSelected}
            },
            getTags() {
                getAllTagApi().then((res) => {
                    for (let i in res.data) {
                        this.tags.push({text: res.data[i], value: i})
                    }
                });
            },
        }
    }
</script>

<style scoped>
    .mut-tags {
        position: relative;
        background: #ffffff;
    }
    .tag {
        margin: 0.2rem;
        background-color: #e8f5fa;
        border-radius: 0.6rem;
        color: #0086cc;
        cursor: pointer;
        display: inline-block;
        font-size: 0.8rem;
        padding: 0 0.6rem;
    }
    .tag.active {
        color: #ff4949;
        background-color:#f8eeee;
        box-shadow: 0 1px 1px -1px #ff4949;
    }

    .tag-box{
        min-width: 500px;
    }
</style>











<!--
<template>
    <div class="mut-tags">
        <div class="scroll">
            <div class="tag-box">
                <span class="tag" :class="tagClass(tag.value)" v-for="(tag, index) in tags"  :key="index" @click="toggleTag(tag.value)">
                    {{ tag.text }}
                </span>
            </div>
        </div>
    </div>
</template>

<script>
    import {getAllTagApi} from '../../api/app'
    import {Tag, Icon} from 'vant';
    export default {
        name: "Tags",
        props: {
            activeId:{
                type: String,
                default: () => {
                    return "0"
                }
            }
        },
        components: {
            [Tag.name]:Tag,
            [Icon.name]:Icon,
        },
        data() {
            return {
                tags:[],
                tagsValue: [],
            }
        },
        created(){
            this.getTags();
            if (this.activeId !== "0") {
                this.toggleTag(this.activeId);
            }
        },
        methods: {
            getTags() {
                getAllTagApi().then((res) => {
                    for (let i in res.data) {
                        this.tags.push({text: res.data[i], value: i})
                    }
                });
            },
            toggleTag(tagId) {
                const index = this.tagsValue.findIndex((t) => {
                    return t === tagId;
                });
                if (index >= 0){
                    this.tagsValue.splice(index, 1);
                }else{
                    this.tagsValue.push(tagId);
                }

                this.$emit("click", this.tagsValue)
            },
            tagClass(tagId) {
                const isSelected = this.tagsValue.includes(tagId);
                return {'active':isSelected}
            }
        }
    }
</script>

<style scoped>
    .mut-tags {
        position: relative;
        background: #ffffff;
    }
    .tag {
        margin: 0.2rem;
        background-color: #e8f5fa;
        border-radius: 0.6rem;
        color: #0086cc;
        cursor: pointer;
        display: inline-block;
        font-size: 0.8rem;
        padding: 0 0.6rem;
    }
    .tag.active {
        color: #ff4949;
        background-color:#f8eeee;
        box-shadow: 0 1px 1px -1px #ff4949;
    }

    .tag-box{
        min-width: 500px;
    }
</style>
-->
