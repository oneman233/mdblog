import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {HomepageComponent} from "./homepage/homepage.component";
import {ArticleComponent} from "./article/article.component";

const routes: Routes = [
  { path: 'articles', component: HomepageComponent },      // 博客首页
  { path: 'article/:id', component: ArticleComponent },    // 文章详情页面
  { path: '', redirectTo: '/articles', pathMatch: 'full' } // 首页重定向
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
