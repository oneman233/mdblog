import { Injectable } from '@angular/core';
import { Article } from 'src/app/article'
import { Observable } from "rxjs";
import { HttpClient } from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class ArticleService {

  private ArticlesURL = "http://localhost:9999/v1/articles"
  private ArticleURL = 'http://localhost:9999/v1/article/'

  constructor(private httpClient: HttpClient) { }

  getArticles(): Observable<Article[]> {
    return this.httpClient.get<Article[]>(this.ArticlesURL)
  }

  getArticle(id: string | null): Observable<Article> {
    return this.httpClient.get<Article>(this.ArticleURL + id)
  }
}
