import { ChangeDetectorRef, Component, SecurityContext } from '@angular/core';
import { AbstractViewComponent } from '../../abstract-view/abstract-view.component';
import { SignpostView, View } from '../../../models/content';
import { parse } from 'marked';
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-signpost',
  templateUrl: './signpost.component.html',
  styleUrls: ['./signpost.component.scss'],
})
export class SignpostComponent extends AbstractViewComponent<SignpostView> {
  message: string;
  component: View;
  trigger: View;

  directions = {
    column: false,
  };

  order = {
    reverse: true,
  };

  constructor(
    private cdr: ChangeDetectorRef,
    private readonly sanitizer: DomSanitizer
  ) {
    super();
  }

  update() {
    this.directions.column = this.v.config.column;
    this.order.reverse = this.v.config.reverse;
    this.component = this.v.config.component;
    this.trigger = this.v.config.trigger;

    const html = parse(this.v.config.message);
    this.message = this.sanitizer.sanitize(SecurityContext.HTML, html);
    this.cdr.markForCheck();
  }
}
