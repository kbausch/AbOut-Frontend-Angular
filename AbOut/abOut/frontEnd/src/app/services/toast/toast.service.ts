import { Injectable, TemplateRef } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ToastService {

  private toasts: any[] = [];

  constructor() { }

  // The public method for showing an error.
  public show(textTitle: string, textOrTpl: string | TemplateRef<any>, options: any = {}) {
    this.toasts.push({ textTitle, textOrTpl, ...options });
  }

  // The public method for removing the toast.
  public remove(toast) {
    this.toasts = this.toasts.filter(t => t !== toast);
  }
}
