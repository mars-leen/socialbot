<template>
    <a-spin :spinning="spinning" size="large">
        <div class="crawler-item">
            <crawler-card v-for="c in crawlerItem" :key="c.Id" :media="c" :category-list="category"></crawler-card>
            <a-back-top :visibilityHeight="10"  @click="top"/>
        </div>
    </a-spin>
</template>

<script>
    import {Button, Icon, Modal, Form, Input, List, Row, Col, Select, Tag, Spin, BackTop} from 'ant-design-vue'
    import CrawlerCard from "../../components/crawler/Card"
    import {
        listRandCrawlerItemApi,
    } from "../../api/robotCrawler"
    import {listCategoryWithTagsApi} from "../../api/category"

    export default {
        name: "CrawlerItem",
        components: {
            CrawlerCard,
            [List.name]: List,
            [Spin.name]: Spin,
            [Tag.name]: Tag,
            [BackTop.name]: BackTop,

            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
            [Row.name]: Row,
            [Col.name]: Col,

            [List.Item.name]: List.Item,
            [List.Item.Meta.name]: List.Item.Meta,
            [Button.name]: Button,
            [Icon.name]: Icon,
            [Modal.name]: Modal,
            [Form.name]: Form,
            [Form.Item.name]: Form.Item,
            [Input.name]: Input,
            [Input.TextArea.name]: Input.TextArea,
        },
        created() {
            this.listCrawlerItem();
            this.listCategory();
           /* this.listCrawlerItemConfig();*/
        },
        data() {
            return {
                spinning:false,
                crawlerItem: [],
                category:[],
            }
        },
        methods: {
            listCrawlerItem(){
                this.spinning = true;
                listRandCrawlerItemApi().then(res => {
                    this.spinning = false;
                    if (!res) {
                        return
                    }
                    this.crawlerItem = res.data
                })
            },
            listCategory() {
                listCategoryWithTagsApi().then(res => {
                    if (!res) {
                        return
                    }
                    this.category = res.data;
                })
            },
            top(){
                this.listCrawlerItem()
            },
        }
    }
</script>

<style scoped>
    .crawler-item {min-height: 350px;width: 100%;}
</style>
