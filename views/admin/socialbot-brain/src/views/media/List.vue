<template>
    <div class="Tag">
        <content-item>

            <template slot="header">
                <a-select mode="default" style="width: 200px;margin-right: 5px;" v-model="activeCategory">
                    <a-select-option v-for="(cate,index) in category" :key="index" :value="cate.Id">{{cate.Title}}</a-select-option>
                </a-select>
            </template>
            <div class="list" slot="body">
                <a-list itemLayout="horizontal" :dataSource="mediaList.items"  :grid="{ gutter: 16, xs: 1, sm: 1, md: 2, lg: 3, xl: 3, xxl: 3 }">
                    <a-list-item slot="renderItem" slot-scope="item,index" >
                        <a-card >
                            <router-link class="action" slot="extra" :to="'/dashboard/media/edit/' + item.Uri"><a-icon  type="edit" /></router-link>
                            <div slot="cover" class="cover-box">
                                <img  alt="cover" :src="item.Cover" class="cover"/>
                            </div>
                            <a-card-meta class="desc" :description="item.Title"></a-card-meta>
                        </a-card>
                    </a-list-item>
                    <div v-if="!mediaList.finished" slot="loadMore" :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }">
                        <a-spin v-if="mediaList.loading" /><a-button v-else @click="listMedias()">loading more</a-button>
                    </div>
                </a-list>
            </div>
        </content-item>
    </div>
</template>

<script>
    import {Button, Icon, Card,Spin, List,Select} from 'ant-design-vue'
    import ContentItem from '../../components/content/Content'
    import {listCategoryApi} from "../../api/category"
    import {listMediasApi} from "../../api/media"
    export default {
        name: "MediaList",
        components: {
            ContentItem,
            [List.name]: List,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
            [List.Item.name]: List.Item,
            [List.Item.Meta.name]: List.Item.Meta,
            [Button.name]: Button,
            [Icon.name]: Icon,
            [Card.name]: Card,
            [Card.Meta.name]: Card.Meta,
            [Spin.name]: Spin,
        },
        created(){
            this.listCategory();
            this.listMedias()
        },
        data() {
            return {
                category:[],
                activeSort:0,
                activeCategory:0,
                mediaList: {
                    items: [],
                    lastId: 0,
                    loading: true,
                    finished: true
                }
            }
        },
        methods:{
            listMedias(){
                this.mediaList.loading= true;
                this.mediaList.finished= false;
                listMediasApi({lastId:this.mediaList.lastId, category:this.activeCategory, sort:this.activeSort}).then(res => {
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
                        this.mediaList.lastId = res.data[l - 1].LastId
                    }
                    for (let i = 0; i < l; i++) {
                        this.mediaList.items.push(res.data[i]);
                    }
                })
            },
            listCategory(){
                this.loading = true;
                listCategoryApi().then(res => {
                    this.loading = false;
                    if (!res) {
                        return
                    }
                    this.category = res.data;
                    this.category.splice(0,0,{
                        Id:0,
                        Title:"全部分类"
                    });
                })
            },
            reloadList(){
                this.mediaList.items =[];
                this.mediaList.lastId= 0;
                this.listMedias()
            }
        },
        watch: {
            activeCategory() {
                this.reloadList()
            }
        }
    }
</script>

<style scoped>
    .cover-box {
        max-height: 200px;
        overflow: hidden;
    }
    .cover-box .cover {
        object-fit: contain;
        width: 100%;
        height:100%
    }
    .action {
        cursor: pointer;
        margin-left: 6px;
    }
</style>
