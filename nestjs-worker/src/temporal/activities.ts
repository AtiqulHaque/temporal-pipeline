import { Logger } from '@nestjs/common';
import { InvoiceService } from '../invoice/invoice.service';
import { createInvoiceActivity } from '../activities/createInvoice.activities';
import { formatMessageActivity } from '../activities/formatMessage.activities';
import { CreateInvoiceInput } from '../activities/createInvoice.types';

const logger = new Logger('TemporalActivities');

export function createActivities(invoiceService: InvoiceService) {
  return {
    formatMessage: async (greeting: string): Promise<string> => {
      logger.log(`[NestJS] formatMessage activity started for: ${greeting}`);
      return formatMessageActivity(greeting);
    },

    createInvoice: async (
      input: CreateInvoiceInput,
    ): Promise<Awaited<ReturnType<typeof createInvoiceActivity>>> => {
      logger.log(
        `[NestJS] createInvoice activity started for: ${input.customerName}`,
      );
      return createInvoiceActivity(invoiceService, input);
    },
  };
}
