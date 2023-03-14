import { Component, OnInit } from '@angular/core';
import { NzModalService } from 'ng-zorro-antd/modal';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { NewPaymentLinksComponent } from 'src/app/modals/new-payment-link/new-payment-link.component';
import {
  FunctionsService,
  FunctionsServiceInterface,
} from 'src/app/services/functions.service';

@Component({
  selector: 'app-payment-links',
  templateUrl: './payment-links.component.html',
  styleUrls: ['./payment-links.component.less'],
})
export class PaymentLinksComponent implements OnInit {
  constructor(
    private functionsService: FunctionsService,
    private notify: NzNotificationService,
    private modal: NzModalService
  ) {}

  public links!: FunctionsServiceInterface.Receive.listLink[];

  public listLink() {
    this.functionsService.listLink().subscribe(
      (response) => {
        response.forEach((link) => {
          link.expand = false;
        });

        this.links = response;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public newLink() {
    this.modal
      .create({
        nzTitle: 'New Payment Link',
        nzContent: NewPaymentLinksComponent,
        nzFooter: null,
        nzWidth: '50%',
      })
      .afterClose.subscribe(() => {
        this.listLink();
      });
  }

  ngOnInit(): void {
    this.listLink();
  }
}
