<template>
    <div class="card">
        <div class="card-header">
            <a-button :loading="delLoading" class="btn" type="danger" @click="delContent">删除</a-button>
        </div>
        <div class="card-body">
            <masonry :cols="{default: 3, 1200: 3, 768: 3, 576: 1}" :gutter="{default: '1rem'}" >
                <div v-for="(v, i) in media.Medias" :key="i" class="card-img" :class="imgClass(v.Download)"
                     @click="toggleImg(v.Download)">
                   <img  :src="v.Show" alt="img" class="relative-img">
                </div>
            </masonry>
        </div>
        <div class="card-bottom">
            <div class="title">
                {{media.Title}}
            </div>
            <div class="recommend">
                <a-checkbox @change="recommend" :default-checked="activeRecommend">是否推荐</a-checkbox>
            </div>
            <div class="cate">
                <a-radio-group v-model="activeCate">
                    <a-radio v-for="c in categoryList" :key="c.Id" :value="c.Id">{{c.ShortName}}</a-radio>
                </a-radio-group>
            </div>

            <div class="tag">
                <span :class="tagClass(tag.Id)" v-for="tag in tagList()" :key="tag.Id" @click="toggleTag(tag)">
                    {{tag.ShortName}}
                </span>
            </div>
            <div class="title">
                <a-select style="width: 100%"
                          showSearch
                          :value="search.value"
                          placeholder="搜索"
                          :defaultActiveFirstOption="false"
                          :showArrow="false"
                          :filterOption="false"
                          @search="searchCopywriter"
                          @change="searchCopywriterChange"
                          notFoundContent="not found">
                    <a-select-option v-for="d in search.data" :key="d.Id">{{d.Title}}</a-select-option>
                </a-select>
                <a-textarea v-model="activeTitle"  placeholder="Basic usage"/>
            </div>

            <div class="submit">
                <a-button :loading="delLoading" class="btn" type="danger" @click="delContent">删除</a-button>
                <a-button :loading="pubLoading" class="btn" @click="pubContent">发布</a-button>
            </div>
        </div>
        <div v-show="isDelEd" class="pop"><span class="deleted">已删除</span></div>
        <div v-show="isPubEd" class="pop"><span class="deleted">已发布</span></div>
    </div>
</template>

<script>
    import Vue from 'vue'
    import VueMasonry from 'vue-masonry-css'
    import {Button, Radio, Tag, Input, Checkbox, Select} from 'ant-design-vue'
    import {deleteCrawlerItemApi, addSocialMediaFromCrawlerApi} from "../../api/robotCrawler";
    import {searchCopywriterApi} from "../../api/copywriter";
    Vue.use(VueMasonry);

    export default {
        name: "CrawlerCard",
        props: {
            categoryList: {
                type: Array,
                default: () => {
                    return []
                }
            },
            media: {
                type: Object,
                default: () => {
                    return {
                        Id: 0,
                        Title: "",
                        Medias: []
                    }
                }
            },
        },
        components: {
            [Button.name]: Button,
            [Input.name]: Input,
            [Input.TextArea.name]: Input.TextArea,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
            [Radio.name]: Radio,
            [Checkbox.name]: Checkbox,
            [Radio.Group.name]: Radio.Group,
            [Tag.name]: Tag,
            [Tag.CheckableTag.name]: Tag.CheckableTag,
        },
        data() {
            return {
                delLoading: false,
                pubLoading: false,

                isDelEd: false,
                isPubEd: false,
                search:{
                    value:undefined,
                    data:[],
                },

                activeTitle:"",
                activeCate: 1,
                activeRecommend: false,
                activeTags: [],
                activeUrl: [],
            }
        },
        methods: {
            toggleTag(tag) {
                const index = this.activeTags.findIndex((t) => {
                    return t === tag.Id;
                });
                if (index >= 0) {
                    this.activeTags.splice(index, 1);
                } else {
                    this.activeTags.push(tag.Id);
                }
            },
            tagClass(tagId) {
                let isSelected = this.activeTags.includes(tagId);
                return {'active': isSelected}
            },
            delContent() {
                if (this.delLoading) {
                    return
                }
                this.delLoading = true;
                // eslint-disable-next-line no-console
                console.log(this.media.Id);
                const f = new FormData;
                f.append("id", this.media.Id);
                deleteCrawlerItemApi(f).then((res) => {
                    this.delLoading = false;
                    if (!res) {
                        return
                    }
                    this.isDelEd = true
                })
            },
            pubContent() {
                if (this.pubLoading) {
                    return
                }
                this.pubLoading = true;
                const medias = [];
                for (let i in this.activeUrl) {
                    medias[i] = {url:this.activeUrl[i]}
                }
                const f = new FormData;
                f.append("title", this.activeTitle);
                f.append("medias", JSON.stringify(medias));
                f.append("tags", this.activeTags);
                f.append("cid", this.activeCate);
                f.append("needFetch", true);
                f.append("recommend", this.activeRecommend);
                f.append("crawlerItemId", this.media.Id);
                console.log(this.media.Id);
                addSocialMediaFromCrawlerApi(f).then((res) => {
                    this.pubLoading = false;
                    if (!res) {
                        return
                    }
                    this.isPubEd = true
                })

            },
            toggleImg(url) {
                const index = this.activeUrl.findIndex((u) => {
                    return u === url;
                });
                if (index >= 0) {
                    this.activeUrl.splice(index, 1);
                } else {
                    this.activeUrl.push(url);
                }
            },
            imgClass(url) {
                let isSelected = this.activeUrl.includes(url);
                return {'active': isSelected}
            },
            tagList() {
                if (this.categoryList.length === 0){
                    return
                }
                const index = this.categoryList.findIndex((cate) => {
                    return cate.Id === this.activeCate;
                });
                let nowCate = this.categoryList[index];
                return nowCate.Tags
            },
            searchCopywriter(value) {
                searchCopywriterApi({key:value}).then(res=>{
                    if (!res) {
                        return
                    }
                    this.search.data = res.data
                })
            },
            searchCopywriterChange(value){
               const index = this.search.data.findIndex((d)=>{
                   return d.Id === value
               });
                this.search.value = this.search.data[index].Title;
                this.activeTitle = this.search.data[index].Description;
            },
            recommend(e){
                this.activeRecommend = e.target.checked
            }
        }
    }
</script>

<style scoped>
    .card {
        background: #ffffff;
        margin-bottom: 1rem;
        position: relative;
    }

    .card-img {
        width: 100%;
        position: relative;
        border: 2px solid #ffffff;
    }

    .card-img.active {
        border: 2px solid #ff6a1e;
    }

    .card-bottom {
        padding: 0.6rem;
    }

    .card-header {
        text-align: right;
        padding: 0.4rem
    }

    .card-bottom div {
        margin-bottom: 0.5rem;
    }

    .card .submit {
        text-align: center;
    }

    .card .submit .btn {
        margin-left: 0.5rem;
    }

    .tag span {
        display: inline-block;
        cursor: pointer;
        margin-right: 0.3rem;
        background: #f0f0f0;
        border-radius: 0.8rem;
        padding: 0 0.5rem;
        font-size: 0.8rem;
        margin-bottom: 0.2rem;
    }

    .tag span.active {
        color: #ff6a1e;
    }

    .pop {
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        background: #ff4e3f;
        opacity: 0.5;
    }

    .deleted {
        position: absolute;
        top: 50%;
        background: #ff4e3f;
        opacity: 0.8;
        font-size: 8rem;
        color: #ffffff;
        text-align: center;
        display: inline-block;
        width: 100%;
    }

    .relative-img {
        font-size: 0;
        vertical-align: top;
        line-height: 0;
        display: block;
        width: 100%;
    }
</style>
