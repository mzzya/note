// tslint:disable:no-any
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { format, compareAsc } from 'date-fns'

const count = 5;

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {
  constructor(private http: HttpClient, private msg: NzMessageService, private notification: NzNotificationService) { }
  // tslint:disable-next-line:no-any
  data: any[] = [];
  pageIndex = 1;
  pageSize = 5;
  total = 0;
  addLoading = false;
  dataLoading = false;

  ngOnInit(): void {
    this.loadData(this.pageIndex, this.pageSize);
  }

  getData(pi: number, ps: number, callback: (res: any) => void): void {

    this.http.get("http://localhost:8081/v1/goods?PageIndex=" + pi + "&PageSize=" + ps).subscribe((res: any) => callback(res));
  }

  loadData(pi: number, ps: number): void {
    this.dataLoading = true;
    setTimeout(() => {
      this.dataLoading = false;
    }, 500);
    this.getData(pi, ps, (res: any) => {
      this.data = res.Data;
      this.total = res.Total;
      this.notification.info(
        '总条数',
        res.Total, {
          nzDuration: 1000,
        }
      );
    })
    // this.data = new Array(5).fill({}).map((_, index) => {
    //   return {
    //     href: 'http://ant.design',
    //     title: `ant design part ${index} (page: ${pi})`,
    //     avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
    //     description: 'Ant Design, a design language for background applications, is refined by Ant UED Team.',
    //     content:
    //       'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.'
    //   };
    // });
  }

  timeFormat(date: number): string {
    return format(date * 1000, 'yyyy-MM-dd HH:mm:ss');
  }
  pageIndexChange(pi: number): void {
    this.pageIndex = pi;
    this.loadData(this.pageIndex, this.pageSize);
  }
  pageSizeChange(ps: number): void {
    this.pageIndex = 1;
    this.pageSize = ps;
    this.loadData(this.pageIndex, this.pageSize);
  }
  add(): void {
    this.addLoading = true;
    setTimeout(() => {
      this.addLoading = false;
    }, 300);
    let g = {
      ID: Math.round(Date.now()),
      Name: "aaa",
      Bn: "bbb",
      Price: 10,
      Pic: "http://img.alicdn.com/imgextra/i1/735276822/TB2.bqVjfxNTKJjy0FjXXX6yVXa_!!735276822-2-beehive-scenes.png",
      Content: "河马男朋友抱枕靠枕床头靠垫大靠背睡觉枕头大号床上长条枕可爱女",
      CreateTime: Math.round(Date.now() / 1000),
      UpdateTime: 0,
    }
    this.http.post("http://localhost:8081/v1/goods", g, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      }
    }).subscribe((res: any) => {
      if (res.Status === true) {
        this.notification.success(
          '添加成功',
          g.ID.toString(), {
            nzDuration: 1000,
          }
        );
      } else {
        this.notification.error(
          '添加失败',
          JSON.stringify(res),
          {
            nzDuration: 3000,
          }
        );
      }
      this.loadData(1, this.pageSize);
    });
  }
}
