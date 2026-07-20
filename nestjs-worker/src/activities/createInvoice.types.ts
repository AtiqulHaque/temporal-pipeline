export interface CreateInvoiceInput {
  customerName: string;
  amount: number;
  currency?: string;
  workflowId?: string;
  invoiceNumber?: string;
}

export interface CreateInvoiceResult {
  id: string;
  invoiceNumber: string;
  customerName: string;
  amount: string;
  currency: string;
  status: string;
  workflowId: string | null;
}
