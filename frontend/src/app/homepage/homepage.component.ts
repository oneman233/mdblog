import { Component, OnInit } from '@angular/core';
import { ArticleService } from "../services/article/article.service";
import { Article } from "../article";

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent implements OnInit {

  articles: Article[] = [];

  constructor(
    private articleService: ArticleService
  ) { }

  ngOnInit(): void {
    this.articleService.getArticles()
      .subscribe(articles => this.articles = articles);
  }

}
