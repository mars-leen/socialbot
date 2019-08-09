<template>
    <div>
        <nav-bar :is-show-nav="false"></nav-bar>
        <div class="container">
            <tags class="tags"  @click="loadTags"></tags>
            <div class="line-label">
                <span class="label">Relation Posts</span>
            </div>
            <van-list v-model="mediaList.loading" :finished="mediaList.finished" finished-text="no more" @load="onLoad">
                <masonry :cols="{default: 3, 1204: 2, 768: 2, 576: 1}" :gutter="{default: '15px'}">
                    <card v-for="(media, i) in mediaList.items" :key="i" :keyIndex="i" :media="media"></card>
                </masonry>
            </van-list>
        </div>
    </div>

</template>
<script>
    import {List} from 'vant';
    import {getListByTagsApi} from '../../api/media'
    import Tags from '../../components/media/Tags'
    import Card from '../../components/media/Card'
    import NavBar from "../../components/navbar/NavBar"
    import {mapGetters, mapMutations} from "vuex"

    export default {
        data() {
            return {
                sort:0,
                isLoading: false,
                mediaList: {
                    items: [],
                    lastId: "0",
                    loading: false,
                    finished: false
                }
            }
        },
        components: {
            // eslint-disable-next-line vue/no-unused-components
            NavBar, Tags, Card, [List.name]: List,
        },
        methods: {
            ...mapGetters({
                getSelectedTags: "getSelectedTags"
            }),
            ...mapMutations({
                updateSelectedTags : "updateSelectedTags"
            }),
            loadTags(){
                this.mediaList.items = [];
                this.mediaList.lastId = 0;
                this.mediaList.loading = false;
                this.mediaList.finished = false;
                window.scrollTo(0, 10);
            },
            onLoad() {
                getListByTagsApi(this.mediaList.lastId, this.getSelectedTags().join(","), this.sort).then((res) => {
                    // eslint-disable-next-line no-console
                    this.mediaList.loading = false;
                    if (!res) {
                        this.mediaList.finished = true;
                        return
                    }
                    let l = res.data.length;
                    if (l < 15) {
                        this.mediaList.finished = true;
                    }
                    if (l !== 0) {
                        this.mediaList.lastId = res.data[l -1].LastId
                    }
                    for (let i = 0; i < l; i++) {
                        this.mediaList.items.push(res.data[i]);
                    }
                })
            },
        }
    }
</script>

<style scoped>
    .tags {padding: 1rem 0 0;}
    @media (max-width: 576px) {
        .tags {
            padding:0.8rem 0.8rem 0 0.8rem;
        }
    }
</style>