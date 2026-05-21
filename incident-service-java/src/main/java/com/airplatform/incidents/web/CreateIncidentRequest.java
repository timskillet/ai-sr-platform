package com.airplatform.incidents.web;

public record CreateIncidentRequest(String service, String severity, String message) {
}
