import { Component, OnInit } from '@angular/core';
import { Article } from "../article";
import { ActivatedRoute } from "@angular/router";
import { ArticleService } from "../services/article/article.service";

@Component({
  selector: 'app-article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.css']
})
export class ArticleComponent implements OnInit {

  article: Article = {};

  constructor(
    private route: ActivatedRoute,
    private articleService: ArticleService
  ) { }

  ngOnInit(): void {
    // 获取路由参数
    let id = this.route.snapshot.paramMap.get('id');
    this.articleService.getArticle(id)
      .subscribe(article => this.article = article)
  }

}
